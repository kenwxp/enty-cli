package util

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dm20151123 "github.com/alibabacloud-go/dm-20151123/client"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dm20151123.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dm.aliyuncs.com")
	_result = &dm20151123.Client{}
	_result, _err = dm20151123.NewClient(config)
	return _result, _err
}

func _main(email, checkCode string) (_err error) {
	client, _err := CreateClient(tea.String("LTAI7zF9r542KCuW"), tea.String("KKNOLUkQaxWXZOCPx1nghq3Bh8TMOf"))
	if _err != nil {
		return _err
	}

	singleSendMailRequest := &dm20151123.SingleSendMailRequest{
		AccountName:    tea.String("manager@investors.entysquare.com"),
		AddressType:    tea.Int32(1),
		ToAddress:      tea.String(email),
		Subject:        tea.String("Filer验证码"),
		ReplyToAddress: tea.Bool(false),
		ReplyAddress:   tea.String(""),
		HtmlBody:       tea.String("邮箱验证码: " + checkCode + "，请勿泄露"),
		TextBody:       tea.String(""),
		FromAlias:      tea.String("Filer"),
	}
	// 复制代码运行请自行打印 API 的返回值
	_, _err = client.SingleSendMail(singleSendMailRequest)
	if _err != nil {
		return _err
	}
	return _err
}

func SingleMail(email, checkCode string) {
	err := _main(email, checkCode)
	if err != nil {
		fmt.Print(err)
	}
}
