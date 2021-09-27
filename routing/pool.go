package routing

import (
	"entysquare/filer-backend/chain/filecoin"
	"entysquare/filer-backend/jsonerror"
	"entysquare/filer-backend/storage"
	"entysquare/filer-backend/storage/types"
	"entysquare/filer-backend/util"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func getUserFilData(
	req *http.Request, db *storage.Database, filerId int64,
) util.JSONResponse {
	ctx := req.Context()

	balanceIncomeList, err := db.SelectFilerBalanceIncomeByFilerId(ctx, nil, strconv.FormatInt(filerId, 10))
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.NotFound("filerAccountIncome can not find successfully"),
		}
	}

	balanceIncome := types.FilerBalanceIncome{
		FilerId:              strconv.FormatInt(filerId, 10),
		PledgeSum:            "0",
		HoldPower:            "0",
		TotalIncome:          "0",
		FreezeIncome:         "0",
		Balance:              "0",
		TotalAvailableIncome: "0",
	}
	if len(balanceIncomeList) > 0 {
		balanceIncome = balanceIncomeList[0]
	}

	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: UserFilResponse{
			ProfitAll:       util.StrNanoFILToFilStr(balanceIncome.TotalIncome, "4"),
			ProfitBalance:   util.StrNanoFILToFilStr(balanceIncome.Balance, "4"),
			ProfitAvailable: util.StrNanoFILToFilStr(balanceIncome.TotalAvailableIncome, "4"),
			ProfitLock:      util.StrNanoFILToFilStr(balanceIncome.FreezeIncome, "4"),
			ProfitToday:     util.StrNanoFILToFilStr(balanceIncome.DayRaiseIncome, "4"),
			PledgeAll:       util.StrNanoFILToFilStr(balanceIncome.PledgeSum, "4"),
			Power:           balanceIncome.HoldPower,
		},
	}

}
func getHtFilNodeData(
	req *http.Request, db *storage.Database,
) util.JSONResponse {
	//ctx := req.Context()
	//htPool, err := db.SelectHtFilerPool(ctx, nil)
	//fmt.Println(htPool, err)
	//if err != nil || htPool == (types.FilerPool{}) {
	//	return util.JSONResponse{
	//		Code: http.StatusForbidden,
	//		JSON: jsonerror.NotFound("filerPool can not find successfully"),
	//	}
	//}
	fmt.Println("api: filecoin.FilNodeUrlData~")
	fmt.Println(filecoin.FilNodeUrlData)
	var qualityPower float64          //字节单位
	var blockRewardAll float64        //NanoFIL 单位
	var blockReward24 float64         //NanoFIL 单位
	var miningEfficiencyFloat float64 //FIL 单位
	for _, v := range filecoin.FilNodeUrlData {
		qualityPower += v.QualityPower
		blockRewardAll += v.BlockRewardAll
		blockReward24 += v.BlockReward24
		miningEfficiencyFloat += v.MiningEfficiencyFloat
	}
	size, typeStr := util.FileSizeStr(qualityPower)
	blockRewardAllStr := util.StrNanoFILToFilStr8(fmt.Sprintf("%.4f", blockRewardAll))
	blockReward24Str := util.StrNanoFILToFilStr8(fmt.Sprintf("%.4f", blockReward24))
	miningEfficiencyFloatStr := fmt.Sprintf("%.4f", miningEfficiencyFloat/float64(len(filecoin.FilNodeUrlData)))

	//size,typeStr := util.FileSizeStr(qualityPower)
	//poolIncome := db.SelectFilerPoolIncomeByNodeIdLast(ctx, htPool.NodeId)
	//avg := util.CalculateString(poolIncome.TotalIncome, poolIncome.TotalPower, "div")
	blocks, _ := db.SelectFilerBlockAndTimeByTime(req.Context(), time.Now().AddDate(0, 0, -30), time.Now())
	var list []BlockResponse
	for _, v := range blocks {
		list = append(list, BlockResponse{
			Date:     v.Time,
			BlockNum: util.StrNanoFILToFilStr(v.Num, "4"),
		})
	}

	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: HtFilNodeResponse{
			PowerAll:      fmt.Sprintf("%.2f", size) + " " + typeStr, //默认返回 T 单位
			Produce24h:    blockReward24Str,
			Produce24hAvg: miningEfficiencyFloatStr,
			ProduceAll:    blockRewardAllStr, //Nano单位换算
			Blocks:        list,
		},
	}

}
