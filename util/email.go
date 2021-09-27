package util

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"strconv"
)

type ErrMail struct {
	MailTo  []string //收件人
	Subject string   //标题
	Body    string   //正文
}

func SendMailCode(mail string, code string) {
	//定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码

	//mailConn := map[string]string{
	//  "user": "xxx@163.com",
	//  "pass": "your password",
	//  "host": "smtp.163.com",
	//  "port": "465",
	//}
	//	"user": "506765498@qq.com",
	//	"pass": "gyynaavwfgaubjbi",
	mailConn := map[string]string{
		"user": "506765498@qq.com",
		"pass": "gyynaavwfgaubjbi",
		"host": "smtp.qq.com",
		"port": "587",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "filtab")) //这种方式可以添加别名，即“XX官方”
	//说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mail)                                //发送给多个用户
	m.SetHeader("Subject", "test 邮箱验证码")                   //设置邮件主题
	m.SetBody("text/html", "<p>邮箱验证码：</p><p>"+code+"</p>") //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println("send fail", "err = ", err)
		return
	}
	fmt.Println("send successfully")
}
