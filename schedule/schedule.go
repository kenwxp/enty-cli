package schedule

import (
	"context"
	"database/sql"
	"entysquare/enty-cli/chain"
	"entysquare/enty-cli/storage"
	"entysquare/enty-cli/storage/types"
	"entysquare/enty-cli/util"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type StatType int

const (
	TaskBlockStat StatType = iota
	TaskOrderStat
	TaskAccountStat
	TaskPoolStat
	TaskNodeStat
	TaskBalanceStat
	TaskMain
)

var statTypeStr = [...]string{
	"TASK_BLOCK_STAT",
	"TASK_ORDER_STAT",
	"TASK_ACCOUNT_STAT",
	"TASK_POOL_STAT",
	"TASK_NODE_STAT",
	"TASK_BALANCE_STAT",
	"TASK_MAIN",
}

func (d StatType) String() string {
	return statTypeStr[d]
}

func StatisticTask(db *storage.Database) {
	// 1. run chain sync serve
	go AutoUpdateBlockTable(db)

	// 2. run schedule of daily batch process
	go AutoScheduleRun(db)

	// 3. set up ht fil node dataset auto inject
	//go filecoin.ScanMiner(db)

	//4. set FilNodeUrlPowers line chart data
	// go chain.SetFilNodeUrlPowersServer(db)
}

func AutoUpdateBlockTable(db *storage.Database) {
	for {
		chain.FilProfitScanningByFilScout(db)
		time.Sleep(time.Second * 10 * 60)
	}
}

func AutoScheduleRun(db *storage.Database) {
	ch := cronReady()
	for {
		select {
		case <-ch:
			//统计时间为当天执行时最早时间 的前一天
			statTime := util.TimeStringToTime(util.TimeNow().Format("2006-01-02"), "00:00:00", "").AddDate(0, 0, -1)
			//全节点统计
			err := mainControl(db, statTime, "ALLNODE")
			if err != nil {
				fmt.Println(" main statistic Schedule err ==> ["+statTime.Format("2006-01-02")+"]", err)
			}
		}
	}
}

func cronReady() chan interface{} {
	// ch thatnotify batch start
	ch := make(chan interface{})

	//get current timestamp
	nowTs := util.TimeNow().Unix()
	//set next schedule running timestamp (set 00:30:00 run schedule)
	timeSuffix := "00:30:00"
	targetTs := util.TimeStringToTime(util.TimeNow().Format("2006-01-02"), timeSuffix, "").Unix()
	if nowTs > targetTs {
		targetTs = util.TimeStringToTime(util.TimeNow().AddDate(0, 0, 1).Format("2006-01-02"), timeSuffix, "").Unix()
	}

	//get duration between to timestamp and set timer for first schedule
	go func() {
		leftTimeTs := targetTs - nowTs
		leftTime := time.Duration(leftTimeTs * 1000 * 1000 * 1000)
		timer := time.NewTimer(leftTime)
		// get 1 day duration
		dayTime := 24 * 60 * 60 * time.Second
		fmt.Println("left time :", leftTime, " day round time:", dayTime)

		mid := <-timer.C
		ch <- mid
		timer.Stop()
		// count each round as a full day time left to base time
		ticker := time.NewTicker(dayTime)
		go func() {
			for {
				midTick := <-ticker.C
				ch <- midTick
				fmt.Println("tick batch process")
			}
		}()
	}()
	return ch
}

func ManualScheduleRun(db *storage.Database, statTimeStr string, nodeId string) (err error) {
	//获取当前统计时间与预设统计时间相差天数
	durationDays, err := util.GetDuratinDaysFromCurrentStatTime(statTimeStr)
	//遍历日期区间的所有日期的统计
	for i := durationDays; i >= 0; i-- {
		statTime := util.TimeStringToTime(util.TimeNow().AddDate(0, 0, -1-i).Format("2006-01-02"), "00:00:00", "")
		err = mainControl(db, statTime, nodeId)
		if err != nil {
			fmt.Println(" main statistic Schedule err ==> ["+statTime.Format("2006-01-02")+"]", err)
			return err
			//return errors.New("exec mainControl>SelectStatControlInfo failed")
		}
	}
	return nil
}
func mainControl(db *storage.Database, statTime time.Time, nodeId string) (err error) {
	ctx := context.TODO()
	statTimeStr := statTime.Format("2006-01-02")
	//检查是否有指定统计时间和nodeId 的批量在跑stat_state = '1'
	statControl, err := db.SelectStatControlInfo(ctx, nil, TaskMain.String(), statTimeStr, "ALLNODE")
	if err != nil {
		fmt.Println("exec mainControl>SelectStatControlInfo failed:")
		return err
	}
	if statControl != nil && statControl.StatState == "1" {
		fmt.Println("[" + statTimeStr + "] statistic skipped ,because other schedule is running... ")
		return nil
	}
	// 删除统计控制表中当前统计时间之后的控制记录
	if nodeId == "" || nodeId == "ALLNODE" {
		err = db.DeleteAllStatControlFromStatTime(ctx, nil, statTimeStr)
		if err != nil {
			fmt.Println("exec mainControl>DeleteAllStatControlFromStatTime failed")
			return err
		}
	} else {
		//删除当前统计日期之后，指定节点的所有批量控制记录
		err = db.DeleteStatControlFromStatTimeByNodeId(ctx, nil, statTimeStr, nodeId)
		if err != nil {
			fmt.Println("exec mainControl>DeleteStatControlFromStatTimeByNodeId failed")
			return err
		}
		//删除当前统计日期之后，全节点统计的批量控制记录（TASK_MAIN,TASK_BALANCE_STAT)
		err = db.DeleteStatControlFromStatTimeByNodeId(ctx, nil, statTimeStr, "ALLNODE")
		if err != nil {
			fmt.Println("exec mainControl>DeleteStatControlFromStatTimeByNodeId failed")
			return err
		}
	}
	//新增主统计控制记录
	err = insertStatControlToDB(db, ctx, nil, TaskMain, statTimeStr, "ALLNODE")
	if err != nil {
		fmt.Println("exec mainControl>insertStatControlToDB failed:")
		return err
	}
	//事务开始
	err = db.WithTransaction(func(txn *sql.Tx) error {
		//查询全部node节点
		nodeList, err := db.SelectAllFilerPool(ctx, txn)
		if err != nil {
			fmt.Println("exec mainControl>SelectAllFilerPool failed:")
			return err
		}
		for _, node := range nodeList {
			// nodeId设置成特殊值ALLNODE 或 空 时，全节点进行跑批
			if nodeId == "" || nodeId == "ALLNODE" {
				err = scheduleExecute(db, ctx, txn, statTime, node.NodeId)
				if err != nil {
					fmt.Println("exec mainControl>scheduleExecute failed:")
					return err
				}
			} else {
				//只有匹配到节点列表中的节点才能跑
				if nodeId == node.NodeId {
					err = scheduleExecute(db, ctx, txn, statTime, node.NodeId)
					if err != nil {
						fmt.Println("exec mainControl>scheduleExecute failed:")
						return err
					}
				}
			}
		}
		// 全节点统计完成后，进行账户余额平账
		err = statisticBalanceIncome(db, ctx, txn, statTime)
		if err != nil {
			fmt.Println("exec mainControl>statisticBalanceIncome failed:")
			return err
		}
		return nil
	})

	if err != nil {
		if err1 := updateStatControlToDB(db, ctx, nil, TaskMain, statTimeStr, "ALLNODE", "3", err.Error()); err1 != nil {
			fmt.Println("exec mainControl>updateStatControlToDB failed:", err1)
			return err1
		}
		fmt.Println("exec mainControl>db.WithTransaction failed:")
		return err
	}
	//更新TASK_MAIN 状态为2-成功
	err = updateStatControlToDB(db, ctx, nil, TaskMain, statTimeStr, "ALLNODE", "2", "successful")
	if err != nil {
		fmt.Println("exec mainControl>updateStatControlToDB failed:")
		return err
	}
	return nil
}

func scheduleExecute(db *storage.Database, ctx context.Context, txn *sql.Tx, statTime time.Time, nodeId string) (err error) {
	statTimeStr := statTime.Format("2006-01-02")
	fmt.Println("[" + statTimeStr + " " + nodeId + "] scheduleExecute start ==============================")

	//新增TASK_BLOCK_STAT统计控制记录
	err = insertStatControlToDB(db, ctx, nil, TaskNodeStat, statTimeStr, nodeId)
	if err != nil {
		fmt.Println("exec scheduleExecute>insertStatControlToDB failed:")
		return err
	}

	//统计开始
	// statistic block info from table filer_block_temp to filer_block
	err = statisticBlockInfo(db, ctx, txn, statTime, nodeId)
	if err != nil {
		fmt.Println("exec scheduleExecute>statisticBlockInfo failed:")
		return err
	}
	// statistic order income from table filer_block to filer_order_income
	err = statisticOrderIncome(db, ctx, txn, statTime, nodeId)
	if err != nil {
		fmt.Println("exec scheduleExecute>statisticOrderIncome failed:")
		return err
	}
	// statistic account income from table filer_order_income to filer_account_income
	err = statisticAccountIncome(db, ctx, txn, statTime, nodeId)
	if err != nil {
		fmt.Println("exec scheduleExecute>statisticAccountIncome failed:")
		return err
	}
	// statistic pool income from table filer_block to filer_pool_income
	err = statisticPoolIncome(db, ctx, txn, statTime, nodeId)
	if err != nil {
		fmt.Println("exec scheduleExecute>statisticPoolIncome failed:")
		return err
	}
	//更新TASK_MAIN 状态为2-成功
	err = updateStatControlToDB(db, ctx, nil, TaskNodeStat, statTimeStr, nodeId, "2", "successful")
	if err != nil {
		fmt.Println("exec scheduleExecute>updateStatControlToDB failed:")
		return err
	}
	fmt.Println("[" + statTimeStr + " " + nodeId + "] scheduleExecute end ==============================")
	return err
}

func statisticBlockInfo(db *storage.Database, ctx context.Context, txn *sql.Tx, statTime time.Time, nodeId string) (err error) {
	statTimeStr := statTime.Format("2006-01-02")
	fmt.Println("==== [" + statTimeStr + " " + nodeId + "] statisticBlockInfo start =========================")
	//删除filer_block当前节点，当前统计日期之后的记录
	err = db.DeleteFilerBlockIncomeFromStatTime(ctx, txn, statTimeStr, nodeId)
	if err != nil {
		fmt.Println("exec statisticBlockInfo>DeleteFilerBlockFromStatTime failed:")
		return err
	}
	//新增TASK_BLOCK_STAT统计控制记录
	err = insertStatControlToDB(db, ctx, txn, TaskBlockStat, statTimeStr, nodeId)
	if err != nil {
		fmt.Println("exec statisticBlockInfo>insertStatControlToDB failed:")
		return err
	}
	// 获取filer_block_temp 的记录
	blockStatList, err := db.SelectStatisticBlockInfoByNodeId(ctx, txn, statTime, nodeId)
	if err != nil {
		fmt.Println("exec statisticBlockInfo>insertStatControlToDB failed:")
		return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskBlockStat, err)
	}
	fmt.Println("---------- blockStatList:", blockStatList)
	if len(blockStatList) > 0 {
		// 由于filer_block不统计累加项，这里将统计的数据直接新增至filer_block
		for _, filerBlock := range blockStatList {
			fmt.Println("------------- new filerBlock:", filerBlock)
			if err = db.InsertFilerBlockIncome(ctx, txn, &filerBlock); err != nil {
				fmt.Println("exec statisticBlockInfo>InsertFilerBlock failed:")
				return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskBlockStat, err)
			}
		}
	}
	//更新TASK_BLOCK_STAT 状态为2-成功
	err = updateStatControlToDB(db, ctx, txn, TaskBlockStat, statTimeStr, nodeId, "2", "successful")
	if err != nil {
		fmt.Println("exec statisticBlockInfo>updateStatControlToDB failed:")
		return err
	}
	fmt.Println("==== [" + statTimeStr + " " + nodeId + "] statisticBlockInfo end =========================")
	return
}

func statisticOrderIncome(db *storage.Database, ctx context.Context, txn *sql.Tx, statTime time.Time, nodeId string) (err error) {
	//获取统计ID （yyyy-MM-dd)
	statTimeStr := statTime.Format("2006-01-02")
	fmt.Println("==== [" + statTimeStr + " " + nodeId + "] statisticOrderIncome start =========================")
	//删除filer_order_income当前节点，当前统计日期之后的记录
	err = db.DeleteFilerOrderIncomeFromStatTime(ctx, txn, statTime.Format("2006-01-02"), nodeId)
	if err != nil {
		fmt.Println("exec statisticOrderIncome>DeleteFilerOrderIncomeFromStatTime failed:")
		return err
	}

	//新增TASK_BLOCK_STAT统计控制记录
	err = insertStatControlToDB(db, ctx, txn, TaskOrderStat, statTimeStr, nodeId)
	if err != nil {
		fmt.Println("exec statisticOrderIncome>insertStatControlInfo failed:")
		return err
	}
	//校验当前节点，当前统计日期 前置任务 TASK_BLOCK_STAT任务是否成功
	statControl, err := db.SelectStatControlInfo(ctx, txn, TaskBlockStat.String(), statTimeStr, nodeId)
	if err != nil {
		fmt.Println("exec statisticOrderIncome>SelectStatControlInfo failed:")
		return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskOrderStat, err)
	}
	//前置任务失败，则直接失败
	if statControl.StatState == "3" {
		fmt.Println("Pre-task TaskBlockStat failed")
		return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskOrderStat, errors.New("pre-task task_block_stat failed"))
	}
	// 查询所有状态state!=9 的持仓信息列表，并遍历
	orderList, err := db.SelectAllInProgressOrderByNodeId(ctx, txn, statTimeStr, nodeId)
	if err != nil {
		fmt.Println("exec statisticOrderIncome>SelectAllInProgressOrderByNodeId failed:")
		return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskOrderStat, err)
	}
	fmt.Println("---------- orderList:", orderList)
	for _, filerOrder := range orderList {
		fmt.Println("========== [" + statTimeStr + " " + nodeId + " " + filerOrder.FilerId + "] order income start ...")
		//每个迭代中，取持仓信息的产品生效时间，下单时间，周期，服务费和统计日期，进行校验，
		//抽取成方法返回值：ValuedBlockRange.Code = 0-收益未开始，1-开始生效 2-有收益，3-有收益且今日到期 9-收益结算完毕 -1 err
		valuedBlockRange := util.GetValuedBlockRange(filerOrder.OrderTime, filerOrder.ValidPlan, filerOrder.Period, statTimeStr)
		fmt.Println("------------- valuedBlockRange:" + strconv.Itoa(valuedBlockRange.Code) + ", " + valuedBlockRange.Start + ", " + valuedBlockRange.End)

		if valuedBlockRange.Code == 0 {
			//还未生效，不作处理
			fmt.Println("========== [" + statTimeStr + " " + nodeId + " " + filerOrder.FilerId + "] order income end ...")
			continue
		} else if valuedBlockRange.Code == 1 || valuedBlockRange.Code == 2 || valuedBlockRange.Code == 3 {
			//有爆块收益结算
			//特殊处理，今日生效情况
			if valuedBlockRange.Code == 1 {
				//更新持仓状态为2-持仓结束
				err = db.UpdateFilerOrderState(ctx, txn, "1", filerOrder.OrderId)
				if err != nil {
					fmt.Println("exec statisticOrderIncome>UpdateFilerOrderState failed:")
					return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskOrderStat, err)
				}
			}
			//特殊处理，今日到期的情况
			if valuedBlockRange.Code == 3 {
				//更新持仓状态为2-持仓结束
				err = db.UpdateFilerOrderState(ctx, txn, "2", filerOrder.OrderId)
				if err != nil {
					fmt.Println("exec statisticOrderIncome>UpdateFilerOrderState failed:")
					return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskOrderStat, err)
				}
			}
			beginTimeStr := valuedBlockRange.Start
			endTimeStr := valuedBlockRange.End
			//根据统计区间获取爆块列表
			blockList, err := db.SelectFilerBlockIncomeByStatRange(ctx, txn, beginTimeStr, endTimeStr, nodeId)
			if err != nil {
				fmt.Println("exec statisticOrderIncome>SelectFilerBlockByStatRange failed")
				return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskOrderStat, err)
			}
			fmt.Println("------------- blockList:", blockList)
			//获取前一天的持仓平账信息
			preStatTime := statTime.AddDate(0, 0, -1).Format("2006-01-02")
			orderIncomeList, err := db.SelectFilerOrderIncomeByStatTime(ctx, txn, preStatTime, filerOrder.OrderId, filerOrder.NodeId)
			if err != nil {
				fmt.Println("exec statisticOrderIncome>SelectFilerOrderIncomeByStatTime failed")
				return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskOrderStat, err)
			}
			fmt.Println("------------- pre orderIncomeList:", preStatTime, orderIncomeList)

			filerOrderIncome := types.FilerOrderIncome{
				OrderId:              filerOrder.OrderId,
				FilerId:              filerOrder.FilerId,
				NodeId:               filerOrder.NodeId,
				TotalIncome:          "0",
				FreezeIncome:         "0",
				TotalAvailableIncome: "0",
				DayAvailableIncome:   "0",
				DayRaiseIncome:       "0",
				DayDirectIncome:      "0",
				DayReleaseIncome:     "0",
				StatTime:             statTimeStr,
			}
			if len(orderIncomeList) > 0 {
				filerOrderIncome = orderIncomeList[0]
			}
			//总量数据需要取上一天的
			statTotalIncome := filerOrderIncome.TotalIncome
			statTotalFreeze := filerOrderIncome.FreezeIncome
			statTotalAvailableIncome := filerOrderIncome.TotalAvailableIncome //累计已释放收益 = 昨日累计收益+当日释放收益
			//增量数据直接初始化为0
			statDayAvailableIncome := "0" //	当日释放收益(当日直接25%+往日75%平账 ）（25%）
			statDayRaiseIncome := "0"     //	日产出收益（当日直接25%+当日冻结75%） |性释放收益（往日75%平账）
			statDayDirectIncome := "0"    //  当日直接收益（25%）
			statDayReleaseIncome := "0"   //  当日线性释放收益（往日75%平账）

			//遍历列表，1、爆块时间等于当前统计日期的区块奖励分配25%部分 2、其余情况分配75%部分，并获取总数，
			for _, filerBlock := range blockList {

				income := util.FSToIS(util.CalculateString(filerOrder.HoldPower, filerBlock.GainPerTib, "mulBigF")) //持有算力x每T收益计算
				incomeRate := util.CalculateString("1", filerOrder.ServiceRate, "subBigFU")                         //1-服务费率
				income = util.FSToIS(util.CalculateString(income, incomeRate, "mulBigF"))                           //持仓获得此爆块的总收益

				if filerBlock.StatTime == statTimeStr {
					//1、爆块时间等于当前统计日期的区块奖励分配25%部分
					statTotalIncome = util.FSToIS(util.CalculateString(statTotalIncome, income, "addBigFH"))
					statDayRaiseIncome = income

					freeze := util.FSToIS(util.CalculateString(statDayRaiseIncome, "0.75", "mulBigF"))
					statTotalFreeze = util.FSToIS(util.CalculateString(statTotalFreeze, freeze, "addBigFH"))

					//可用收益进行累加
					statDayDirectIncome = util.FSToIS(util.CalculateString(statDayRaiseIncome, "0.25", "mulBigF"))
					statDayAvailableIncome = util.FSToIS(util.CalculateString(statDayAvailableIncome, statDayDirectIncome, "addBigFH"))
				} else {
					//2、其余情况分配75%部分，并获取总数，
					freeze := util.FSToIS(util.CalculateString(income, "0.75", "mulBigF"))
					divide := util.CalculateString("1", "180", "divBigF")
					freezePart := util.FSToIS(util.CalculateString(freeze, divide, "mulBigF")) //取1/180

					statTotalFreeze = util.FSToIS(util.CalculateString(statTotalFreeze, freezePart, "subBigFH"))
					statDayReleaseIncome = util.FSToIS(util.CalculateString(statDayReleaseIncome, freezePart, "addBigFH"))
					statDayAvailableIncome = util.FSToIS(util.CalculateString(statDayAvailableIncome, freezePart, "addBigFH"))
				}
			}
			filerOrderIncome.StatTime = statTimeStr
			filerOrderIncome.TotalIncome = statTotalIncome
			filerOrderIncome.FreezeIncome = statTotalFreeze
			filerOrderIncome.TotalAvailableIncome = util.FSToIS(util.CalculateString(statTotalAvailableIncome, statDayAvailableIncome, "addBigFH"))
			filerOrderIncome.DayAvailableIncome = statDayAvailableIncome
			filerOrderIncome.DayRaiseIncome = statDayRaiseIncome
			filerOrderIncome.DayDirectIncome = statDayDirectIncome
			filerOrderIncome.DayReleaseIncome = statDayReleaseIncome

			fmt.Println("------------- new filerOrderIncome:", filerOrderIncome)
			err = db.InsertFilerOrderIncome(ctx, txn, &filerOrderIncome)
			if err != nil {
				fmt.Println("exec statisticOrderIncome>InsertFilerOrderIncome failed")
				return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskOrderStat, err)
			}
		} else if valuedBlockRange.Code == 9 {
			//所有收益已结清，更新持仓状态为9
			err = db.UpdateFilerOrderState(ctx, txn, "9", filerOrder.OrderId)
			if err != nil {
				fmt.Println("exec statisticOrderIncome>UpdateFilerOrderState failed")
				return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskOrderStat, err)
			}
		}
		fmt.Println("========== [" + statTimeStr + " " + nodeId + " " + filerOrder.FilerId + "] order income end ...")
	}

	//更新TASK_BLOCK_STAT 状态为2-成功
	err = updateStatControlToDB(db, ctx, txn, TaskOrderStat, statTimeStr, nodeId, "2", "successful")
	if err != nil {
		fmt.Println("exec statisticOrderIncome>updateStatControlToDB failed")
		return err
	}
	fmt.Println("==== [" + statTimeStr + " " + nodeId + "] statisticOrderIncome end =========================")
	return err
}

func statisticAccountIncome(db *storage.Database, ctx context.Context, txn *sql.Tx, statTime time.Time, nodeId string) (err error) {
	//获取统计ID （yyyy-MM-dd)
	statTimeStr := statTime.Format("2006-01-02")
	fmt.Println("==== [" + statTimeStr + " " + nodeId + "] statisticAccountIncome start =========================")
	//删除filer_account_income 当前节点，当前统计日期之后的记录
	err = db.DeleteFilerAccountIncomeFromStatTime(ctx, txn, statTime.Format("2006-01-02"), nodeId)
	if err != nil {
		fmt.Println("exec statisticAccountIncome>DeleteFilerAccountIncomeFromStatTime failed")
		return err
	}

	//新增TASK_BLOCK_STAT统计控制记录
	err = insertStatControlToDB(db, ctx, txn, TaskAccountStat, statTimeStr, nodeId)
	if err != nil {
		fmt.Println("exec statisticAccountIncome>insertStatControlToDB failed")
		return err
	}

	//校验当前节点，当前统计日期 前置任务TASK_ORDER_STAT任务是否成功
	statControl, err := db.SelectStatControlInfo(ctx, txn, TaskOrderStat.String(), statTimeStr, nodeId)
	if err != nil {
		fmt.Println("exec statisticAccountIncome>SelectStatControlInfo failed")
		return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskAccountStat, err)
	}
	//前置任务失败，则直接失败
	if statControl.StatState == "3" {
		fmt.Println("Pre-task TaskOrderStat failed")
		return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskAccountStat, errors.New("pre-task task_order_stat failed"))
	}

	// 根据filerId聚合持仓平账表filer_order_income 收益数据，获得统计列表
	accountIncomeStatList, err := db.SelectStatisticOrderIncomeByNodeId(ctx, txn, statTimeStr, nodeId)
	if err != nil {
		fmt.Println("exec statisticAccountIncome>SelectStatisticOrderIncomeByNodeId failed")
		return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskAccountStat, err)
	}

	fmt.Println("---------- accountIncomeStatList:", accountIncomeStatList)
	// 新增至filer_block
	for _, accountIncomeStat := range accountIncomeStatList {
		fmt.Println("========== [" + statTimeStr + " " + nodeId + " " + accountIncomeStat.FilerId + "] account income start ...")
		//根据filerId 和前一天的统计时间去获取filer_account_income 中存量的数据
		preStatTime := statTime.AddDate(0, 0, -1).Format("2006-01-02")
		accountIncomeList, err := db.SelectFilerAccountIncomeByStatTime(ctx, txn, preStatTime, accountIncomeStat.FilerId, nodeId)
		if err != nil {
			fmt.Println("exec statisticAccountIncome>SelectFilerAccountIncomeByStatTime failed")
			return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskAccountStat, err)
		}
		fmt.Println("------------- pre accountIncomeList:", accountIncomeList)
		filerAccountIncome := types.FilerAccountIncome{
			FilerId:              accountIncomeStat.FilerId,
			NodeId:               nodeId,
			PledgeSum:            "0",
			HoldPower:            "0",
			TotalIncome:          "0",
			FreezeIncome:         "0",
			TotalAvailableIncome: "0",
			DayAvailableIncome:   "0",
			DayRaiseIncome:       "0",
			DayDirectIncome:      "0",
			DayReleaseIncome:     "0",
			StatTime:             statTimeStr,
		}
		if len(accountIncomeList) > 0 {
			filerAccountIncome = accountIncomeList[0]
		}
		//将统计信息更新到存量数据中，并在filer_account_income新增当天数据
		//累计项信息根据统计信息直接替换
		filerAccountIncome.PledgeSum = accountIncomeStat.PledgeSum
		filerAccountIncome.HoldPower = accountIncomeStat.HoldPower
		filerAccountIncome.TotalIncome = util.FSToIS(accountIncomeStat.TotalIncome)
		filerAccountIncome.FreezeIncome = util.FSToIS(accountIncomeStat.FreezeIncome)
		filerAccountIncome.TotalAvailableIncome = util.FSToIS(accountIncomeStat.TotalAvailableIncome)
		filerAccountIncome.DayAvailableIncome = util.FSToIS(accountIncomeStat.DayAvailableIncome)
		filerAccountIncome.DayRaiseIncome = util.FSToIS(accountIncomeStat.DayRaiseIncome)
		filerAccountIncome.DayDirectIncome = util.FSToIS(accountIncomeStat.DayDirectIncome)
		filerAccountIncome.DayReleaseIncome = util.FSToIS(accountIncomeStat.DayReleaseIncome)
		//更新统计时间
		filerAccountIncome.StatTime = statTimeStr

		fmt.Println("------------- new filerAccountIncome:", filerAccountIncome)
		//新增一条新纪录
		err = db.InsertFilerAccountIncome(ctx, txn, &filerAccountIncome)
		if err != nil {
			fmt.Println("exec statisticAccountIncome>InsertFilerAccountIncome failed")
			return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskAccountStat, err)
		}
		fmt.Println("========== [" + statTimeStr + " " + nodeId + " " + accountIncomeStat.FilerId + "] account income end ...")
	}

	//更新TASK_BLOCK_STAT 状态为2-成功
	err = updateStatControlToDB(db, ctx, txn, TaskAccountStat, statTimeStr, nodeId, "2", "successful")
	if err != nil {
		fmt.Println("exec statisticAccountIncome>updateStatControlToDB failed")
		return err
	}
	fmt.Println("==== [" + statTimeStr + " " + nodeId + "] statisticAccountIncome end =========================")
	return
}

func statisticBalanceIncome(db *storage.Database, ctx context.Context, txn *sql.Tx, statTime time.Time) (err error) {
	//获取统计ID （yyyy-MM-dd)
	statTimeStr := statTime.Format("2006-01-02")
	fmt.Println("[" + statTimeStr + "] statisticBalanceIncome start =========================")
	//删除filer_account_income 当前节点，当前统计日期之后的记录
	err = db.DeleteFilerBalanceIncomeFromStatTime(ctx, txn, statTime.Format("2006-01-02"))
	if err != nil {
		fmt.Println("exec statisticBalanceIncome>DeleteFilerBalanceIncomeFromStatTime failed")
		return err
	}

	//新增TASK_BLOCK_STAT统计控制记录
	err = insertStatControlToDB(db, ctx, txn, TaskBalanceStat, statTimeStr, "ALLNODE")
	if err != nil {
		fmt.Println("exec statisticBalanceIncome>insertStatControlToDB failed")
		return err
	}

	//校验当前统计日期 所有节点的前置任务TASK_ACCOUNT_STAT任务是否成功
	statControlList, err := db.SelectAllNodeStatControlInfo(ctx, txn, TaskAccountStat.String(), statTimeStr)
	if err != nil {
		fmt.Println("exec statisticBalanceIncome>SelectAllNodeStatControlInfo failed")
		return handleStatExcption(db, ctx, txn, statTime, "ALLNODE", TaskBalanceStat, err)
	}
	for _, statControl := range statControlList {
		//前置任务失败，则直接失败
		if statControl.StatState == "3" {
			fmt.Println("Pre-task TaskAccountStat failed")
			return handleStatExcption(db, ctx, txn, statTime, "ALLNODE", TaskBalanceStat, errors.New("pre-task task_account_stat failed"))
		}
	}

	// 根据filerId聚合持仓平账表filer_account_income 收益数据，获得统计列表
	balanceIncomeStatList, err := db.SelectStatisticBalanceIncome(ctx, txn, statTimeStr)
	if err != nil {
		fmt.Println("exec statisticBalanceIncome>SelectStatisticBalanceIncome failed")
		return handleStatExcption(db, ctx, txn, statTime, "ALLNODE", TaskBalanceStat, err)
	}

	fmt.Println("---------- balanceIncomeStatList:", balanceIncomeStatList)
	// 新增至filer_block
	for _, balanceIncomeStat := range balanceIncomeStatList {
		fmt.Println("========== [" + statTimeStr + " " + balanceIncomeStat.FilerId + "] account income start ...")
		//根据filerId 和前一天的统计时间去获取filer_account_income 中存量的数据
		preStatTime := statTime.AddDate(0, 0, -1).Format("2006-01-02")
		balanceIncomeList, err := db.SelectFilerBalanceIncomeByStatTime(ctx, txn, preStatTime, balanceIncomeStat.FilerId)
		if err != nil {
			fmt.Println("exec statisticBalanceIncome>SelectFilerBalanceIncomeByStatTime failed")
			return handleStatExcption(db, ctx, txn, statTime, "ALLNODE", TaskBalanceStat, err)
		}
		fmt.Println("------------- pre balanceIncomeList:", balanceIncomeList)
		filerBalanceIncome := types.FilerBalanceIncome{
			FilerId:              balanceIncomeStat.FilerId,
			PledgeSum:            "0",
			HoldPower:            "0",
			TotalIncome:          "0",
			FreezeIncome:         "0",
			Balance:              "0",
			TotalAvailableIncome: "0",
			DayAvailableIncome:   "0",
			DayRaiseIncome:       "0",
			DayDirectIncome:      "0",
			DayReleaseIncome:     "0",
			StatTime:             statTimeStr,
		}
		if len(balanceIncomeList) > 0 {
			filerBalanceIncome = balanceIncomeList[0]
		}
		//将统计信息更新到存量数据中，并在filer_account_income新增当天数据
		//累计项信息根据统计信息直接替换
		filerBalanceIncome.PledgeSum = balanceIncomeStat.PledgeSum
		filerBalanceIncome.HoldPower = balanceIncomeStat.HoldPower
		filerBalanceIncome.TotalIncome = util.FSToIS(balanceIncomeStat.TotalIncome)
		filerBalanceIncome.FreezeIncome = util.FSToIS(balanceIncomeStat.FreezeIncome)
		filerBalanceIncome.TotalAvailableIncome = util.FSToIS(balanceIncomeStat.TotalAvailableIncome)
		filerBalanceIncome.DayAvailableIncome = util.FSToIS(balanceIncomeStat.DayAvailableIncome)
		filerBalanceIncome.DayRaiseIncome = util.FSToIS(balanceIncomeStat.DayRaiseIncome)
		filerBalanceIncome.DayDirectIncome = util.FSToIS(balanceIncomeStat.DayDirectIncome)
		filerBalanceIncome.DayReleaseIncome = util.FSToIS(balanceIncomeStat.DayReleaseIncome)

		///计算余额
		//1、先将当天统计的可用收益增量增加进前一天的余额
		statBalance := util.FSToIS(util.CalculateString(filerBalanceIncome.Balance, balanceIncomeStat.DayAvailableIncome, "addBigFH"))
		//2、可用收益变动需要与余额变动表进行关联计算
		balanceFlowList, err := db.SelectStatisticBalanceFlowListByFilerId(ctx, txn, balanceIncomeStat.FilerId, statTimeStr)
		if err != nil {
			fmt.Println("exec statisticBalanceIncome>SelectStatisticBalanceFlowListByFilerId failed")
			return handleStatExcption(db, ctx, txn, statTime, "ALLNODE", TaskBalanceStat, err)
		}
		//balanceFlow.OperType  0-收益入账 1-手工入账 2-提现出账
		for _, balanceFlow := range balanceFlowList {
			if balanceFlow.OperType == "1" {
				//手工入账的金额重新加到可用收益余额
				statBalance = util.FSToIS(util.CalculateString(statBalance, balanceFlow.Amount, "addBigFH"))
			} else if balanceFlow.OperType == "2" {
				//提现出账的金额重新从可用收益余额中去除
				statBalance = util.FSToIS(util.CalculateString(statBalance, balanceFlow.Amount, "subBigFH"))
			} else if balanceFlow.OperType == "0" {
				//重跑时会出现收益入账这条记录，需要手动删除统计当天的记录
				err = db.DeleteBalanceFlowByOperType(ctx, txn, balanceIncomeStat.FilerId, statTimeStr, "0")
				if err != nil {
					fmt.Println("exec statisticBalanceIncome>DeleteBalanceFlowByOperType failed")
					return handleStatExcption(db, ctx, txn, statTime, "ALLNODE", TaskBalanceStat, err)
				}
			}
		}
		filerBalanceIncome.Balance = statBalance

		//将当天释放收益，新增至余额变动信息表，创建时间设置为统计当天最晚时间
		insertBalanceFlow := types.FilerBalanceFlow{
			FilerId:    balanceIncomeStat.FilerId,
			OperType:   "0",
			Amount:     balanceIncomeStat.DayAvailableIncome,
			CreateTime: strconv.FormatInt(util.TimeStringToTime(statTimeStr, "23:59:59", "").Unix(), 10),
		}
		fmt.Println("------------- new insertBalanceFlow:", insertBalanceFlow)
		err = db.InsertBalanceFlow(ctx, txn, &insertBalanceFlow)
		if err != nil {
			fmt.Println("exec statisticBalanceIncome>InsertBalanceFlow failed")
			return handleStatExcption(db, ctx, txn, statTime, "ALLNODE", TaskBalanceStat, err)
		}

		//更新统计时间
		filerBalanceIncome.StatTime = statTimeStr
		fmt.Println("------------- new filerBalanceIncome:", filerBalanceIncome)
		//新增一条新纪录
		err = db.InsertFilerBalanceIncome(ctx, txn, &filerBalanceIncome)
		if err != nil {
			fmt.Println("exec statisticBalanceIncome>InsertFilerBalanceIncome failed")
			return handleStatExcption(db, ctx, txn, statTime, "ALLNODE", TaskBalanceStat, err)
		}
		fmt.Println("========== [" + statTimeStr + " " + balanceIncomeStat.FilerId + "] account income end ...")
	}

	//更新TASK_BLOCK_STAT 状态为2-成功
	err = updateStatControlToDB(db, ctx, txn, TaskBalanceStat, statTimeStr, "ALLNODE", "2", "successful")
	if err != nil {
		fmt.Println("exec statisticBalanceIncome>updateStatControlToDB failed")
		return err
	}
	fmt.Println("[" + statTimeStr + "] statisticBalanceIncome end =========================")
	return
}

func statisticPoolIncome(db *storage.Database, ctx context.Context, txn *sql.Tx, statTime time.Time, nodeId string) (err error) {
	statTimeStr := statTime.Format("2006-01-02")
	fmt.Println("==== [" + statTimeStr + " " + nodeId + "] statisticPoolIncome start =========================")
	//删除filer_pool_income 当前节点，当前统计日期之后的记录
	err = db.DeleteFilerPoolIncomeFromStatTime(ctx, txn, statTime.Format("2006-01-02"), nodeId)
	if err != nil {
		fmt.Println("exec statisticPoolIncome>DeleteFilerPoolIncomeFromStatTime failed")
		return err
	}

	//新增TASK_POOL_STAT统计控制记录
	err = insertStatControlToDB(db, ctx, txn, TaskPoolStat, statTimeStr, nodeId)
	if err != nil {
		fmt.Println("exec statisticPoolIncome>insertStatControlToDB failed")
		return err
	}

	//校验当前节点，当前统计日期 前置任务TASK_BLOCK_STAT任务是否成功
	statControl, err := db.SelectStatControlInfo(ctx, txn, TaskBlockStat.String(), statTimeStr, nodeId)
	if err != nil {
		fmt.Println("exec statisticPoolIncome>SelectStatControlInfo failed")
		return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskPoolStat, err)
	}
	//前置任务失败，则直接失败
	if statControl.StatState == "3" {
		fmt.Println("Pre-task TaskBlockStat failed")
		return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskPoolStat, errors.New("pre-task task_block_stat failed"))
	}

	var available float64 = 0  //今日可用 25%
	var lock float64 = 0       //今日质押 75%
	var release180 float64 = 0 //180天 线性释放数量
	var num180 int64 = 0       //180天出块总量

	//查询node 180内爆块数据
	blocks, err := db.SelectFilerBlockIncomeByStatRange(
		ctx, txn,
		statTime.AddDate(0, 0, -180).Format("2006-01-02"), //今天前推180天
		statTimeStr, nodeId)
	if err != nil {
		fmt.Println("exec statisticPoolIncome>SelectFilerBlockByStatRange failed")
		return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskPoolStat, err)
	}
	fmt.Println("------------- blockList:", blocks)

	for _, block := range blocks {
		filNum, _ := strconv.ParseInt(block.BlockGain, 10, 64)
		num180 += filNum
		//当天数据
		if block.StatTime == statTimeStr {
			available = float64(filNum) * 0.25
			lock = float64(filNum) * 0.75
			num180 -= filNum
		}
	}
	release180 = (float64(num180) * (1.0 / 180.0)) * 0.75 //总币量的 75% 属于线性释放

	//var dbPool = types.NewFilerPoolIncome();
	addPool, err := db.SelectFilerPoolIncomeByNodeId(ctx, txn,
		nodeId,
		statTime.AddDate(0, 0, -1).Format("2006-01-02"))
	if err != nil {
		fmt.Println("exec statisticPoolIncome>SelectFilerPoolIncomeByNodeId failed")
		return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskPoolStat, err)
	}
	fmt.Println("------------- pre poolIncome:", addPool)
	if addPool == nil {
		addPool = types.NewFilerPoolIncome()
	}
	addPool.NodeId = nodeId
	addPool.Balance = util.FSToIS(util.Addition(addPool.Balance, available, release180))                        // 余额 += 今25% + 线性释放
	addPool.TotalIncome = util.FSToIS(util.Addition(addPool.TotalIncome, available, lock))                      //总收益 += 今天100%
	addPool.AvailableIncome = util.FSToIS(util.Addition(addPool.AvailableIncome, available, release180))        //可用收益 += 今25% + 线性释放
	addPool.FreezeIncome = util.FSToIS(util.Subtraction(util.Addition(addPool.FreezeIncome, lock), release180)) //冻结收益 += 今75% - 线性释放
	addPool.TodayIncomeTotal = util.FSToIS(util.Addition(available, lock))                                      //今日产出收益 = 今天100%
	addPool.TodayIncomeFreeze = util.FSToIS(util.Addition(lock))                                                //今日冻结收益 = 今天75%
	addPool.TodayIncomeAvailable = util.FSToIS(util.Addition(available, release180))                            //今日可用收益 = 今天25% + 线性释放
	addPool.StatTime = statTimeStr                                                                              //统计时间（yyyy-MM-dd）
	if len(blocks) > 0 {
		addPool.TotalPower = blocks[len(blocks)-1].Power //总算力
	} else {
		addPool.TotalPower = "0" //总算力
	}
	fmt.Println("------------- new poolIncome:", addPool)
	err = db.InsertFilerPoolIncome(ctx, txn, addPool)
	if err != nil {
		fmt.Println("exec statisticPoolIncome>InsertFilerPoolIncome failed")
		return handleStatExcption(db, ctx, txn, statTime, nodeId, TaskPoolStat, err)
	}
	//更新TASK_BLOCK_STAT 状态为2-成功
	err = updateStatControlToDB(db, ctx, txn, TaskPoolStat, statTimeStr, nodeId, "2", "successful")
	if err != nil {
		fmt.Println("exec statisticPoolIncome>updateStatControlToDB failed")
		return err
	}
	fmt.Println("==== [" + statTimeStr + " " + nodeId + "] statisticPoolIncome end =========================")
	return
}

//异常处理
func handleStatExcption(db *storage.Database, ctx context.Context, txn *sql.Tx, statTime time.Time, nodeId string, statType StatType, err error) error {
	if err1 := updateStatControlToDB(db, ctx, txn, statType, statTime.Format("2006-01-02"), nodeId, "3", err.Error()); err1 != nil {
		return err1
	}

	if err1 := updateStatControlToDB(db, ctx, nil, TaskNodeStat, statTime.Format("2006-01-02"), nodeId, "3", err.Error()); err1 != nil {
		return err1
	}
	return err
}

func updateStatControlToDB(db *storage.Database, ctx context.Context, txn *sql.Tx, statType StatType, statTimeStr string, nodeId string, statState string, message string) error {
	statControl := initStatControlInfo(statType, statTimeStr, nodeId)
	statControl.StatState = statState
	statControl.Message = message
	err := db.UpdateStatControlInfo(ctx, txn, &statControl)
	return err
}

func insertStatControlToDB(db *storage.Database, ctx context.Context, txn *sql.Tx, statType StatType, statTimeStr string, nodeId string) error {
	// 新增总控记录到批量控制表 并更新至进行时
	statControl := initStatControlInfo(statType, statTimeStr, nodeId)
	statControl.StatState = "1"
	statControl.Message = "running..."
	err := db.InsertStatControlInfo(ctx, txn, &statControl)
	return err
}

func initStatControlInfo(statType StatType, statTimeStr string, nodeId string) types.FilerStatControl {
	insertStatControl := types.FilerStatControl{}
	insertStatControl.StatType = statType.String()
	insertStatControl.StatTime = statTimeStr
	insertStatControl.NodeId = nodeId
	insertStatControl.StatState = "0"
	insertStatControl.Message = "init"
	return insertStatControl
}
