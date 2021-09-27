package chain

import (
	"context"
	"database/sql"
	"entysquare/enty-cli/chain/filecoin"
	"entysquare/enty-cli/storage"
	"entysquare/enty-cli/storage/types"
	"entysquare/enty-cli/util"
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"time"
)

/*
	FIL 扫链记录节点出块数据
*/
func FilProfitScanning() {
	//当前扫高度
	var toHeight int64
	//最新区块高度
	var maxHeight int64

	//获取Database
	db, err := storage.NewDatabase()
	if err != nil {
		panic("[panic!]FilProfitScanning db err 001")
	}

	//获取初始区块高度
	toHeight, err = GetBlockScanningHeight(db)
	if err != nil {
		panic("[panic!]FilProfitScanning db err 001")
	}

	//var address []string
	//address = append(address, "f0154944")
	//address = append(address, "f0109070")
	//address = append(address, "f065609")
	//address = append(address, "f0452222")
	//address = append(address, "f0155212")
	//address = append(address, "f0127573")
	//address = append(address, "f0688083")
	//address = append(address, "f0874295")
	//address = append(address, "f040044")
	//address = append(address, "f0219106")
	nodeList, err := db.SelectFilerPool()
	if err != nil {
		panic("[panic!]filer_Pool db err 001")
	}
	for {
		maxHeight, err = GetBlockHeight()
		//只扫（最新高度-60）不是则休眠10s
		if toHeight < maxHeight-60 {
			var jsonStr string
			var err error

			//获取FIL链 指定高度数据
			err = retry(30, time.Second*5, func() error {
				fmt.Print("retry~ ")
				jsonStr, err = GetBlockHeightData(toHeight)
				return err
			})
			if err != nil {
				panic("[panic!]GetBlockHeightData err :" + err.Error())
			}
			fmt.Println("======================================================")
			fmt.Println("jsonStr;", jsonStr)
			fmt.Println("address:", nodeList)
			fmt.Println("maxHeight:", maxHeight)
			fmt.Println("toHeight:", toHeight)
			//处理包块数据 返回内部包块信息
			list, err := handleBlockMiner(jsonStr, nodeList, "24000000000") //NanoFIL = 10*(-9) FIL，or we can say 1 FIL = 1000000000 NanoFIL

			err = db.WithTransaction(func(txn *sql.Tx) error {
				var ctx = context.TODO()
				for _, filerBlockTemp := range list {
					if err = db.InsertFilerBlockTemp(ctx, txn, &filerBlockTemp); err != nil {
						return err

					}
				}
				return err
			})
			//入库出错 重扫
			if err != nil {
				continue
			}
			if err == nil {
				toHeight++
			}
		} else {
			time.Sleep(time.Second * 10)
		}
	}

	//开始扫块 遍历爆块者

	//获取需要扫的节点数组

	//判断是否包含

	//是
	//插入数据 区块高度，节点号，爆块数，倍数，算力，链时间，数据跟新时间

}

/*
	www.filscout.com 平台获取全部出块数据
*/
func FilProfitScanningByFilScout(db *storage.Database) {
	var ctx = context.TODO()
	var blockTemps []types.FilerBlockTemp

	//遍历需要处理的节点
	nodes, err := db.SelectFilerPool()
	for _, node := range nodes {
		//node db 记录最高区块高度
		maxHeight, _ := db.SelectFilerBlockTempMaxBlockHeightByNodeId(ctx, node.NodeId)
		//算力记录数据
		powerList, err := GetFilScoutPowerList(node.NodeName)
		if err != nil {
			continue
		}
		var pageIndex = 1
	ERGODIC:
		//node name 浏览器取 出块分页数据
		filScoutData, err := GetFilScoutData("https://api.filscout.com/api/v1/block", node.NodeName, pageIndex)
		if err != nil {
			continue
		}
		//出快 详细数据
		for _, block := range filScoutData.Data {
			//出块数据比记录数据 旧
			if block.Height <= int(maxHeight) {
				filScoutData.PageBool = false
				break
			}
			//出块数据还没处理
			if block.ExactReward == "0" {
				continue
			}
			//出块时间
			unix := util.TimeStrToUnix(block.MineTime)
			if unix == -1 {
				return
			}
			//节点名，出块时间，算力记录 得出出块算力
			power, err := GetFilPowerByMinerDatePower(node.NodeName, unix, powerList)
			//得不出算力数据 整个return 从新干
			if err != nil {
				return
			}
			blockTemps = append(blockTemps, types.FilerBlockTemp{
				NodeId:      node.NodeId,
				BlockNum:    FilToNanoFILToBlockNum(block.ExactReward), //出票数量 2票=48
				BlockGain:   FilToNanoFIL(block.ExactReward),           //24.12873782390871 换算 NanoFIL
				BlockHeight: strconv.Itoa(block.Height),
				ChainTime:   strconv.FormatInt(unix, 10),
				Power:       strconv.FormatInt(power, 10),
			})
			fmt.Println("insert:", node.NodeName,
				" ExactReward:", FilToNanoFIL(block.ExactReward),
				" Height:", strconv.Itoa(block.Height),
				" Power:", power)
		} //for - block
		//time.Sleep(time.Second * 2)
		if filScoutData.PageBool {
			pageIndex++
			goto ERGODIC
		}
	} //for - node

	//倒序插入数据
	err = db.WithTransaction(func(txn *sql.Tx) (err error) {
		for i := len(blockTemps) - 1; i > -1; i-- {
			err = db.InsertFilerBlockTemp(ctx, txn, &blockTemps[i])
			if err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		panic("wrong process from chain")
	}
}

/*
	插入全局变量数据
*/
func SetFilNodeUrlPowers(nodes []types.FilerPool) []filecoin.MinerDatePower {

	var list []filecoin.MinerDatePower
	list = append(list, filecoin.MinerDatePower{})
	list = append(list, filecoin.MinerDatePower{})
	list = append(list, filecoin.MinerDatePower{})
	list = append(list, filecoin.MinerDatePower{})
	list = append(list, filecoin.MinerDatePower{})
	for _, node := range nodes {
		fmt.Println(node.NodeName)
		json30d, err := GetFilScoutPowerStatsData(node.NodeName, "30d")
		//fmt.Println(json30d)
		//fmt.Println(json30d)
		if err != nil {
			return list
		}
		powers := gjson.Get(json30d, "data.powers").Array()
		listIndex := 0
		for i := len(powers) - 1; i > -1; i-- {
			if listIndex == 5 {
				break
			}
			fmt.Println(powers[i])
			list[listIndex] = filecoin.MinerDatePower{
				Unix:  powers[i].Get("heightTime").Int(),
				Date:  powers[i].Get("heightTimeStr").String()[0:10],
				Power: list[listIndex].Power + powers[i].Get("rawPower").Int(),
				//PowerStr:  list[i].PowerStr+v.Get("rawPowerStr").String(),
			}
			listIndex++
		}
	}
	return list
}

func SetFilNodeUrlPowersServer(db *storage.Database) {
	for {
		list, err := db.SelectFilerPool()
		if err != nil {
			fmt.Println("SetFilNodeUrlPowersServer err", err)
			time.Sleep(time.Hour)
			continue
		}
		fmt.Println("api SetFilNodeUrlPowersServer list:", list)
		filecoin.FilNodeUrlPowers = SetFilNodeUrlPowers(list)
		time.Sleep(time.Hour)
	}
}

func FilToNanoFIL(num string) string {
	float, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return ""
	}
	return strconv.FormatFloat(float*1000000000.0, 'f', 0, 64)
}
func FilToNanoFILToBlockNum(num string) string {
	float, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return "1"
	}
	if float > 25.0 {
		return "2"
	}
	return "1"
}

func retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			return s.error
		}
		if attempts--; attempts > 0 {
			fmt.Println("retry func error:", err.Error())
			time.Sleep(sleep)
			return retry(attempts, 2*sleep, fn)
		}
		return err
	}
	return nil
}

type stop struct{ error }
