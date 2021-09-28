package service

import (
	"context"
	"database/sql"
	"entysquare/enty-cli/storage"
	"entysquare/enty-cli/storage/types"
	"entysquare/enty-cli/util"
	"errors"
	"fmt"
	"strconv"
	"time"
)

func InsertOrder(db *storage.Database, name string, amount string) error {
	ctx := context.TODO()
	filerId, err := GetFilerIdByName(db, ctx, name)
	if err != nil {
		println("InsertOrder => GetFilerIdByName => err:", err)
		return err
	}
	nanoAmount := util.FSToIS(util.CalculateString(amount, "1000000000", "mulBigF")) // 转化为nanoFIl
	holdPower := util.CalculateString(nanoAmount, "5160000000", "divBigF")           // 设置算力售价为5.16Fil/T
	createTime := strconv.FormatInt(util.TimeNow().Unix(), 10)
	updateTime := strconv.FormatInt(util.TimeNow().Unix(), 10)
	validTime := strconv.FormatInt(util.TimeNow().AddDate(0, 0, 1).Unix(), 10) //设置T+1生效
	EndTime := strconv.FormatInt(util.TimeNow().AddDate(0, 0, 540).Unix(), 10) //结束日期为T+1+(540-1)
	orderInfo := types.FilerOrder{
		FilerId:    filerId,
		PayFlow:    "",
		ProductId:  "91fb8ea2-d435-4709-b933-1f7057b7f9ef", //使用确定产品下单
		HoldPower:  holdPower,
		PayAmount:  nanoAmount,
		OrderTime:  createTime,
		UpdateTime: updateTime,
		ValidTime:  validTime,
		EndTime:    EndTime,
		OrderState: "0",
	}
	err = db.InsertFilerOrder(ctx, nil, &orderInfo)
	if err != nil {
		return err
	}
	return nil
}

func Withdraw(db *storage.Database, name string, amount string) error {
	ctx := context.TODO()
	filerId, err := GetFilerIdByName(db, ctx, name)
	if err != nil {
		println("Withdraw => GetFilerIdByName => err:", err)
		return err
	}
	err = db.WithTransaction(func(txn *sql.Tx) error {

		balanceIncomeList, err := db.SelectFilerBalanceIncomeByFilerId(ctx, txn, filerId) //获取最新的余额变动情况
		if err != nil {
			fmt.Println("Withdraw =>SelectFilerBalanceIncomeByFilerId => err:", err)
			return err
		}
		if len(balanceIncomeList) > 1 {
			fmt.Println("no balance record exist")
			return err
		}
		filerBalanceIncome := balanceIncomeList[0]
		curBalance := filerBalanceIncome.Balance
		nanoAmount := util.FSToIS(util.CalculateString(amount, "1000000000", "mulBigF")) // 转化为nanoFIl
		newBalance := util.FSToIS(util.CalculateString(curBalance, nanoAmount, "subBigFH"))
		filerBalanceIncome.Balance = newBalance
		//余额变动发生在当天，不能去修改前一天的平账数据，否者影响次日批量任务，故而需新增当天的平账记录，用于实时展示，次日批量任务将根据流水记录重新平账出最新的余额
		todayTimeStr := util.TimeNow().Format("2006-01-02")
		statTime := filerBalanceIncome.StatTime
		if todayTimeStr == statTime {
			//查出最新数据的统计时间已为当天，说明有其他余额变动操作，只需更新即可
			err = db.UpdateFilerBalanceIncomeByUuid(ctx, txn, &filerBalanceIncome)
			if err != nil {
				fmt.Println("Withdraw =>UpdateFilerBalanceIncomeByUuid => err:", err)
				return err
			}
		} else {
			//当天余额未变更，新增当天平账信息
			filerBalanceIncome.StatTime = todayTimeStr
			err = db.InsertFilerBalanceIncome(ctx, txn, &filerBalanceIncome)
			if err != nil {
				fmt.Println("Withdraw =>InsertFilerBalanceIncome => err:", err)
				return err
			}
		}
		//新增提币操作流水
		filerBalanceFlow := types.FilerBalanceFlow{
			FilerId:  filerId,
			OperType: "2",        // 操作类型 0-收益入账 1-手工入账 2-提现出账
			Amount:   nanoAmount, // 金额
		}
		db.InsertBalanceFlow(ctx, txn, &filerBalanceFlow)
		if err != nil {
			fmt.Println("Withdraw =>InsertBalanceFlow => err:", err)
			return err
		}
		return nil

	})

	if err != nil {
		return err
	}
	return nil
}
func Deposit(db *storage.Database, name string, amount string) error {
	ctx := context.TODO()
	filerId, err := GetFilerIdByName(db, ctx, name)
	if err != nil {
		println("Deposit => GetFilerIdByName => err:", err)
		return err
	}
	err = db.WithTransaction(func(txn *sql.Tx) error {

		balanceIncomeList, err := db.SelectFilerBalanceIncomeByFilerId(ctx, txn, filerId) //获取最新的余额变动情况
		if err != nil {
			fmt.Println("Deposit =>SelectFilerBalanceIncomeByFilerId => err:", err)
			return err
		}
		if len(balanceIncomeList) > 1 {
			fmt.Println("no balance record exist")
			return err
		}
		filerBalanceIncome := balanceIncomeList[0]
		curBalance := filerBalanceIncome.Balance
		nanoAmount := util.FSToIS(util.CalculateString(amount, "1000000000", "mulBigF")) // 转化为nanoFIl
		newBalance := util.FSToIS(util.CalculateString(curBalance, nanoAmount, "addBigFH"))
		filerBalanceIncome.Balance = newBalance
		//余额变动发生在当天，不能去修改前一天的平账数据，否者影响次日批量任务，故而需新增当天的平账记录，用于实时展示，次日批量任务将根据流水记录重新平账出最新的余额
		todayTimeStr := util.TimeNow().Format("2006-01-02")
		statTime := filerBalanceIncome.StatTime
		if todayTimeStr == statTime {
			//查出最新数据的统计时间已为当天，说明有其他余额变动操作，只需更新即可
			err = db.UpdateFilerBalanceIncomeByUuid(ctx, txn, &filerBalanceIncome)
			if err != nil {
				fmt.Println("Deposit =>UpdateFilerBalanceIncomeByUuid => err:", err)
				return err
			}
		} else {
			//当天余额未变更，新增当天平账信息
			filerBalanceIncome.StatTime = todayTimeStr
			err = db.InsertFilerBalanceIncome(ctx, txn, &filerBalanceIncome)
			if err != nil {
				fmt.Println("Deposit =>InsertFilerBalanceIncome => err:", err)
				return err
			}
		}
		//新增提币操作流水
		filerBalanceFlow := types.FilerBalanceFlow{
			FilerId:  filerId,
			OperType: "1",        // 操作类型 0-收益入账 1-手工入账 2-提现出账
			Amount:   nanoAmount, // 金额
		}
		db.InsertBalanceFlow(ctx, txn, &filerBalanceFlow)
		if err != nil {
			fmt.Println("Deposit =>InsertBalanceFlow => err:", err)
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func QueryIncomeList(db *storage.Database) error {
	ctx := context.TODO()
	balanceList, err := db.SelectLatestBalanceListForEachFiler(ctx, nil)
	if err != nil {
		fmt.Println("QueryIncomeList =>SelectLatestBalanceListForEachFiler => err:", err)
		return err
	}
	if len(balanceList) > 0 {
		fmt.Println("|\t用户名\t|\t质押金额\t|\t持有算力\t|\t总收益\t|\t已释放收益\t|\t可用余额\t|\t修改时间\t|")
		for _, balanceInfo := range balanceList {
			pledge := util.StrNanoFILToFilStr(balanceInfo.PledgeSum, "4")
			power := balanceInfo.HoldPower
			totalIncome := util.StrNanoFILToFilStr(balanceInfo.TotalIncome, "4")
			available := util.StrNanoFILToFilStr(balanceInfo.TotalAvailableIncome, "4")
			balance := util.StrNanoFILToFilStr(balanceInfo.Balance, "4")
			updateTs, err := strconv.ParseInt(balanceInfo.UpdateTime, 10, 64)
			if err != nil {
				updateTs = 0
			}
			updateTime := time.Unix(updateTs, 0).Format("2006-01-02 15:04:05")

			fmt.Println("" +
				"|\t" + balanceInfo.FilerName + "\t" +
				"|\t" + pledge + "\t" +
				"|\t" + power + "\t" +
				"|\t" + totalIncome + "\t" +
				"|\t" + available + "\t" +
				"|\t" + balance + "\t" +
				"|\t" + updateTime + "\t|")
		}
	} else {
		fmt.Println(" ------> no result <------", err)
	}
	return nil
}
func GetFilerIdByName(db *storage.Database, ctx context.Context, name string) (string, error) {
	account, err := db.SelectAccountByName(ctx, nil, name)
	if err != nil {
		return "", err
	}
	if account == nil {
		return "", errors.New("user is not exist")
	}
	return account.FilerId, nil
}
