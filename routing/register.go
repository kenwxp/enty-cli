package routing

import (
	"encoding/json"
	"entysquare/filer-backend/jsonerror"
	"entysquare/filer-backend/storage"
	"entysquare/filer-backend/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// .GetFilter implements GET /_matrix/client/r0/user/{userId}/filter/{filterId}
func register(
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

	reqParam := &RegisterRequest{}
	err = json.Unmarshal(body, reqParam)
	fmt.Println(string(body))
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.BadJSON("json input is wrong"),
		}
	}
	// first to see whether it has been registered
	account, err := db.SelectAccountByPayID(ctx, nil, reqParam.PayID)
	if err != nil || account != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.UserInUse("register error"),
		}
	}

	token := util.RandomString(64) + ":" + strconv.FormatInt(time.Now().Unix(), 10)
	err = db.CreateAccount(ctx, nil, reqParam.PayID, token)
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.Forbidden("login error"),
		}
	}

	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: RegisterResponse{
			Token: strings.Split(token, ":")[0],
		},
	}

}
