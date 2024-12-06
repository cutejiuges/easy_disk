package email

import (
	"crypto/tls"
	"fmt"
	"github.com/cutejiuges/disk_back/micro_services/user_server/conf"
	"github.com/jordan-wright/email"
	"net/smtp"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/5 下午11:42
 * @FilePath: email_client
 * @Description: 创建email用于发送验证码
 */

func SendVerifyCode(emailAddr string, content int64) error {
	client := email.NewEmail()
	emailConf := conf.GetConf().Email
	client.From = fmt.Sprintf("Verify Code <%s>", emailConf.Host)
	client.To = []string{emailAddr}
	client.Subject = "注册验证码"
	client.HTML = []byte(fmt.Sprintf("您的验证码为:<h1>%d</h1>", content))
	err := client.SendWithTLS(
		emailConf.TLSAddr,
		smtp.PlainAuth("", emailConf.Host, emailConf.Password, emailConf.ServerName),
		&tls.Config{InsecureSkipVerify: true, ServerName: emailConf.ServerName},
	)
	return err
}
