package routing

import (
	"entysquare/filer-backend/jsonerror"
	"entysquare/filer-backend/storage"
	"entysquare/filer-backend/util"
	"net/http"
	"strconv"
)

// .GetFilter implements GET /_matrix/client/r0/user/{userId}/filter/{filterId}
func getProfitLogList(
	req *http.Request, db *storage.Database, filerId int64,
) util.JSONResponse {
	ctx := req.Context()
	maps, err := db.SelectFilerBalanceIncomeListByFilerId(ctx, nil, strconv.FormatInt(filerId, 10))

	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.NotFound("filerBalanceIncome can not find successfully"),
		}
	}
	var result []ProfitLogResponse
	for _, val := range maps {
		if err != nil {
			return util.JSONResponse{
				Code: http.StatusOK,
				JSON: jsonerror.NotFound("string can not transfer successfully"),
			}
		}
		result = append(result, ProfitLogResponse{
			Uuid:                 val.Uuid,
			PledgeSum:            util.StrNanoFILToFilStr(val.PledgeSum, "4"),
			HoldPower:            val.HoldPower,
			TotalIncome:          util.StrNanoFILToFilStr(val.TotalIncome, "4"),
			FreezeIncome:         util.StrNanoFILToFilStr(val.FreezeIncome, "4"),
			TotalAvailableIncome: util.StrNanoFILToFilStr(val.TotalAvailableIncome, "4"),
			Balance:              util.StrNanoFILToFilStr(val.Balance, "4"),
			DayAvailableIncome:   util.StrNanoFILToFilStr(val.DayAvailableIncome, "4"),
			DayRaiseIncome:       util.StrNanoFILToFilStr(val.DayRaiseIncome, "4"),
			DayDirectIncome:      util.StrNanoFILToFilStr(val.DayDirectIncome, "4"),
			DayReleaseIncome:     util.StrNanoFILToFilStr(val.DayReleaseIncome, "4"),
			StatTime:             val.StatTime,
		})

	}
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: ProfitLogListResponse{
			List: result,
		},
	}

}
