package routing

import (
	"entysquare/filer-backend/httputil"
	"entysquare/filer-backend/storage"
	"entysquare/filer-backend/storage/types"
	"entysquare/filer-backend/util"
	"github.com/gorilla/mux"
	"net/http"
)

// Setup configures the given mux with sync-server listeners

func Setup(
	csMux *mux.Router, accountDB *storage.Database,
) {
	r0mux := csMux.PathPrefix("/r0").Subrouter()

	//登录
	r0mux.Handle("/pool/login",
		httputil.MakeExternalAPI("login", func(req *http.Request) util.JSONResponse {
			return Login(req, accountDB)
		}),
	).Methods(http.MethodPost, http.MethodOptions)

	//注册
	r0mux.Handle("/pool/register",
		httputil.MakeExternalAPI("register", func(req *http.Request) util.JSONResponse {
			return register(req, accountDB)
		}),
	).Methods(http.MethodPost, http.MethodOptions)

	//验证码
	r0mux.Handle("/pool/checkCode",
		httputil.MakeExternalAPI("checkCode", func(req *http.Request) util.JSONResponse {
			return applyCheckCode(req)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
	// /pool/productList 产品列表
	r0mux.Handle("/pool/productList",
		httputil.MakeExternalAPI("productList", func(req *http.Request) util.JSONResponse {
			return getProductList(req, accountDB)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
	// /pool/productone 产品信息
	r0mux.Handle("/pool/productOne",
		httputil.MakeExternalAPI("productOne", func(req *http.Request) util.JSONResponse {
			return getProductInfo(req, accountDB)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
	//orderList 订单列表
	r0mux.Handle("/pool/orderList",
		httputil.MakeAuthAPI("orderList", accountDB, func(req *http.Request, account *types.Account) util.JSONResponse {
			return getOrderList(req, accountDB, account.FilerID)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
	//profitLogList 订单列表
	r0mux.Handle("/pool/profitLogList",
		httputil.MakeAuthAPI("profitLogList", accountDB, func(req *http.Request, account *types.Account) util.JSONResponse {
			return getProfitLogList(req, accountDB, account.FilerID)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
	//矿池数据
	r0mux.Handle("/pool/htFilNodeData",
		httputil.MakeExternalAPI("htFilNodeData", func(req *http.Request) util.JSONResponse {
			return getHtFilNodeData(req, accountDB)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
	//profitLogList 矿池用户数据
	r0mux.Handle("/pool/userFilData",
		httputil.MakeAuthAPI("userFilData", accountDB, func(req *http.Request, account *types.Account) util.JSONResponse {
			return getUserFilData(req, accountDB, account.FilerID)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
}
