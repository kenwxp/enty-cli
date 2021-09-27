package routing

import (
	"encoding/json"
	"entysquare/filer-backend/jsonerror"
	"entysquare/filer-backend/storage"
	"entysquare/filer-backend/util"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type user_data struct {
	User_Name string `json:"user_name"` //邮箱号，手机号...
	Password  string `json:"password"`
	Types     string `json:"types"` //登陆类型，手机验证码，邮箱验证码，谷歌验证
	Code      string `json:"code"`
	Equipment int    `json:"equipment"` //设备参数
}

type UserPreLogin struct {
	UserName   string
	Token      string
	GoogleFlag int
}

// GetFilter implements GET /_matrix/client/r0/user/{userId}/filter/{filterId}
func Login(
	req *http.Request, db *storage.Database,
) util.JSONResponse {
	ctx := req.Context()
	bodyio := req.Body
	body, err := ioutil.ReadAll(bodyio)
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.BadJSON("json input is wrong"),
		}
	}
	reqParam := &LoginRequest{}
	err = json.Unmarshal(body, reqParam)
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.BadJSON("json input is wrong"),
		}
	}
	// first get account
	user, err := db.SelectAccountByPayID(ctx, nil, reqParam.PayID)
	if err != nil || user == nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.Forbidden("user is not exist"),
		}
	}
	token := util.RandomString(64) + ":" + strconv.FormatInt(time.Now().Unix(), 10)
	err = db.UpdateAccountToken(ctx, nil, user.PayID, token)
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.Forbidden("update token error"),
		}
	}
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: LoginResponse{
			Token: strings.Split(token, ":")[0],
		},
	}
}
