package routing

import (
	"encoding/json"
	"entysquare/filer-backend/jsonerror"
	"entysquare/filer-backend/storage"
	"entysquare/filer-backend/util"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getProductList(
	req *http.Request, db *storage.Database,
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
	reqParam := &ProductListRequest{}
	err = json.Unmarshal(body, reqParam)
	fmt.Println(string(body))
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.BadJSON("json input is wrong"),
		}
	}
	// first to see whether it has been registered
	maps, err := db.SelectProductListByState(ctx, nil, reqParam.State)
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.NotFound("product can not find successfully"),
		}
	}
	var result []ProductResponse
	for _, val := range maps {
		//pid, err := strconv.Atoi(val.ProductId)
		if err != nil {
			return util.JSONResponse{
				Code: http.StatusOK,
				JSON: jsonerror.NotFound("string can not transfer successfully"),
			}
		}
		fmt.Println("ProductResponse:", val)
		result = append(result, ProductResponse{
			ProductID:   val.ProductId,
			Currency:    val.CurId,
			Name:        val.ProductName,
			PricePerT:   util.StrNanoFILToFilStr(val.Price, "2"),
			PledgeLimit: util.StrNanoFILToFilStr(val.PledgeMax, "2"),
			CycleDay:    val.Period,
			State:       val.ProductState,
		})

	}
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: ProductListResponse{
			List: result,
		},
	}

}
func getProductInfo(
	req *http.Request, db *storage.Database,
) util.JSONResponse {
	bodyIo := req.Body
	ctx := req.Context()
	body, err := ioutil.ReadAll(bodyIo)

	fmt.Println(string(body))
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusOK,
			JSON: jsonerror.NotFound("io can not read successfully"),
		}
	}

	reqParam := &ProductInfoRequest{}
	err = json.Unmarshal(body, reqParam)
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.BadJSON("json input is wrong"),
		}
	}
	product, err := db.SelectProductInfoById(ctx, nil, reqParam.ProductID)
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.NotFound("product can not find successfully"),
		}
	}

	//pid, err := strconv.Atoi(product.ProductId)
	//if err != nil {
	//	return util.JSONResponse{
	//		Code: http.StatusOK,
	//		JSON: jsonerror.NotFound("string can not transfer successfully"),
	//	}
	//}
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: ProductInfoResponse{
			ProductID:      product.ProductId,
			Currency:       product.CurId,
			State:          product.ProductState,
			Name:           product.ProductName,
			PriceService:   util.Multiplication(product.ServiceRate, 100),
			Note1:          product.Note1,
			Note2:          product.Note2,
			CycleDay:       product.Period,
			GuessProfitDay: "0.0315", //预计日产出
			PricePerT:      util.StrNanoFILToFilStr(product.Price, "2"),
		},
	}

}
