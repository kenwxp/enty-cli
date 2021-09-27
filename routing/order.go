package routing

import (
	"encoding/json"
	"entysquare/filer-backend/jsonerror"
	"entysquare/filer-backend/storage"
	"entysquare/filer-backend/storage/types"
	"entysquare/filer-backend/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// .GetFilter implements GET /_matrix/client/r0/user/{userId}/filter/{filterId}
func getOrderList(
	req *http.Request, db *storage.Database, filerId int64,
) util.JSONResponse {
	bodyIo := req.Body
	ctx := req.Context()
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusOK,
			JSON: jsonerror.NotFound("io can not read successfully"),
		}
	}
	reqParam := &OrderListRequest{}
	err = json.Unmarshal(body, reqParam)
	fmt.Println(string(body))
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.BadJSON("json input is wrong"),
		}
	}
	var filerOrderShowList []types.FilerOrderShow
	if reqParam.State == 0 {
		//未生效订单
		filerOrderShowList, err = db.SelectInvalidOrderListByFilerId(ctx, nil, strconv.FormatInt(filerId, 10))
		for _, filerOrderShow := range filerOrderShowList {
			filerOrderShow.TotalIncome = "0"
			filerOrderShow.TotalAvailableIncome = "0"
			filerOrderShow.FreezeIncome = "0"
			filerOrderShow.StatTime = ""
			filerOrderShow.DayAvailableIncome = "0"
			filerOrderShow.DayRaiseIncome = "0"
			filerOrderShow.DayDirectIncome = "0"
			filerOrderShow.DayReleaseIncome = "0"
		}
	} else {
		filerOrderShowList, err = db.SelectFilerOrderListWithIncomeInfoByFilerIdAndOrderState(ctx, nil, strconv.FormatInt(filerId, 10), strconv.Itoa(reqParam.State))
	}
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.NotFound("order can not find successfully"),
		}
	}
	var result []OrderResponse
	for _, filerOrderShow := range filerOrderShowList {
		result = append(result, OrderResponse{
			OrderId:              filerOrderShow.OrderId,                                                          //订单id
			PayFlow:              filerOrderShow.PayFlow,                                                          //支付流水号
			HoldPower:            filerOrderShow.HoldPower,                                                        //持有算力
			PayAmount:            util.StrNanoFILToFilStr(filerOrderShow.PayAmount, "4"),                          //金额质押
			OrderTime:            util.ConvertTimestampToTimeStr(filerOrderShow.OrderTime, "2006-01-02 15:04:05"), //下单时间 时间戳
			ValidTime:            util.ConvertTimestampToTimeStr(filerOrderShow.ValidTime, "2006-01-02 15:04:05"), //生效时间 时间戳
			OrderState:           filerOrderShow.OrderState,                                                       //订单状态 0-待生效 1-已生效 2-持仓已结束 9-收益结算完成
			ProductName:          filerOrderShow.ProductName,                                                      //产品名
			Period:               filerOrderShow.Period,                                                           //周期
			ValidDays:            filerOrderShow.ValidDays,                                                        //已生效天数
			TotalIncome:          util.StrNanoFILToFilStr(filerOrderShow.TotalIncome, "4"),                        //累计收益
			TotalAvailableIncome: util.StrNanoFILToFilStr(filerOrderShow.TotalAvailableIncome, "4"),               //累计已释放收益
			FreezeIncome:         util.StrNanoFILToFilStr(filerOrderShow.FreezeIncome, "4"),                       //待释放收益
			StatTime:             filerOrderShow.StatTime,                                                         //统计日期 yyyy-mm-dd
			DayAvailableIncome:   util.StrNanoFILToFilStr(filerOrderShow.DayAvailableIncome, "4"),                 //当日释放
			DayRaiseIncome:       util.StrNanoFILToFilStr(filerOrderShow.DayRaiseIncome, "4"),                     //当日产出
			DayDirectIncome:      util.StrNanoFILToFilStr(filerOrderShow.DayDirectIncome, "4"),                    //当日直接释放
			DayReleaseIncome:     util.StrNanoFILToFilStr(filerOrderShow.DayReleaseIncome, "4"),                   //当日线性释放
		})
	}
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: OrderListResponse{
			List: result,
		},
	}

}
