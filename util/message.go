package util

import (
	"fmt"
	client "github.com/yunpian/yunpian-go-sdk/sdk"
)

func SendCheckCodeMessage(mobile, checkcode string) {
	// 发送短信
	ypClient := client.New("ec557d72a53ef29f0aa0c39e79d59814")
	param := client.NewParam(2)
	param[client.MOBILE] = mobile
	param[client.TEXT] = "【 TreasureBox】欢迎使用Filer，您的手机验证码是" + checkcode + "。本条信息无需回复"
	r := ypClient.Sms().SingleSend(param)
	if r.Code != 0 {
		fmt.Print(r.Msg)
	}
}
