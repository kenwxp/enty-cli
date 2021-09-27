package routing

import (
	"encoding/json"
	"entysquare/filer-backend/jsonerror"
	this "entysquare/filer-backend/routing/cache"
	"entysquare/filer-backend/util"
	"fmt"
	client "github.com/yunpian/yunpian-go-sdk/sdk"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CheckCode struct {
	Phone_Num string `json:"phone_num"`
	Email     string `json:"email"`
	Code      int    `json:"code"`
	TimeStamp int    `json:"timestamp"`
}

// GetFilter implements GET /_matrix/client/r0/user/{userId}/filter/{filterId}
func applyCheckCode(
	req *http.Request,
) util.JSONResponse {
	bodyio := req.Body
	body, ioerr := ioutil.ReadAll(bodyio)
	if ioerr != nil {
		return util.JSONResponse{
			Code: http.StatusOK,
			JSON: jsonerror.NotFound("io can not read successfully"),
		}
	}
	var jsonList CheckCode
	marshalErr := json.Unmarshal(body, &jsonList)
	if marshalErr != nil {
		return util.JSONResponse{
			Code: http.StatusOK,
			JSON: jsonerror.NotFound("json can not read successfully"),
		}
	}
	if !util.VerifyEmailFormat(jsonList.Email) && !util.VerifyMobileFormat(jsonList.Phone_Num) {
		return util.JSONResponse{
			Code: http.StatusOK,
			JSON: jsonerror.NotFound("username can not read successfully"),
		}
	}
	Map := make(map[string]string)
	checkCodeCache := this.CacheAll
	checkCode := fmt.Sprintf("%06v", rand.New(rand.NewSource(util.TimeNow().UnixNano())).Int31n(1000000))
	timeStamp := fmt.Sprintf("%d", util.TimeNow().UnixNano()/1000000)

	// 发送短信
	ypClient := client.New("ec557d72a53ef29f0aa0c39e79d59814")
	param := client.NewParam(2)
	param[client.MOBILE] = jsonList.Phone_Num
	param[client.TEXT] = "【 TreasureBox】欢迎使用Investors，您的手机验证码是" + checkCode + "。本条信息无需回复"
	r := ypClient.Sms().SingleSend(param)
	if r.Code != 0 {
		fmt.Print(r.Msg)
	}
	if !strings.EqualFold(jsonList.Phone_Num, "") {
		Map["Phone_Num"] = jsonList.Phone_Num
		Map["CheckCode"] = checkCode
		Map["TimeStamp"] = timeStamp
		checkCodeCache.Set(jsonList.Phone_Num+"phoneMap", Map, 60000*time.Millisecond)
	}
	if !strings.EqualFold(jsonList.Email, "") {
		//util.SendMailCode(jsonList.Email, checkCode)
		util.SingleMail(jsonList.Email, checkCode)
		Map["Email"] = jsonList.Email
		Map["CheckCode"] = checkCode
		Map["TimeStamp"] = timeStamp
		//		fmt.Println(jsonList.Email+"emailMap")
		checkCodeCache.Set(jsonList.Email+"emailMap", Map, 60000*time.Millisecond)
	}
	//ca,_:= checkCodeCache.Get(jsonList.Email+"emailMap")
	//map2 := ca.(map[string] string)
	//s :=map2["CheckCode"]
	//fmt.Print(s)
	jsonList.Code, _ = strconv.Atoi(checkCode)
	jsonList.TimeStamp, _ = strconv.Atoi(timeStamp)
	//tx, txerr := sql.DB.BeginTx()
	//if txerr != nil {
	//	return util.JSONResponse{
	//		Code: http.StatusBadRequest,
	//		JSON: jsonerror.NotFound("tx can not open"),
	//	}
	//}
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: jsonList,
	}

}
