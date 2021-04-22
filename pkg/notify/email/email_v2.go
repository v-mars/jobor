package email

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"jobor/pkg/logger"
)


// MailboxConf 邮箱配置
type MailboxConf struct {
	// 邮件标题
	Title string
	// 邮件内容
	Body string
	// 邮件附件
	Attach   []string
	// 收件人列表
	RecipientList []string
	// 发件人账号
	Sender string
	// 发件人密码，QQ邮箱这里配置授权码
	Password string
	From  string
	// SMTP 服务器地址， QQ邮箱是smtp.qq.com
	SMTPAddr string
	// SMTP端口 QQ邮箱是25
	SMTPPort int
	Tls      bool
}

func NewMail(username, password, SMTPAddr,from string,SMTPPort int, tls bool) MailboxConf {
	return MailboxConf{Sender: username,Password: password, SMTPAddr: SMTPAddr,SMTPPort: SMTPPort,From: from,Tls: tls}
}

func (mailConf MailboxConf) Send (title, body string, RecipientList,AttachList []string) error {
	m := gomail.NewMessage()
	m.SetHeader(`From`, mailConf.Sender)
	m.SetHeader(`To`, RecipientList...)
	m.SetHeader(`Subject`, title)
	m.SetBody(`text/html`, body)
	//m.Attach("./Dockerfile") //添加附件
	if len(AttachList)>0{
		for _,v:=range AttachList{m.Attach(v)}
	}

	dialer := gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.Password)
	//dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	//if mailConf.Tls{dialer.TLSConfig = &tls.Config{InsecureSkipVerify: false}}
	err := dialer.DialAndSend(m)
	if err != nil {
		err = fmt.Errorf("Send Email Fail, %s\n", err.Error())
		logger.Error(err)
		return err
	}
	logger.Debug("Send Email Success")
	return nil
}

func SendMailTest() {
	var mailConf MailboxConf
	mailConf.Title = "测试用gomail发送邮件"
	mailConf.Body = "Good Good Study, Day Day Up!!!!!!"
	mailConf.RecipientList = []string{`xxx@qq.com`}
	mailConf.Sender = `xx@xx.cn`
	mailConf.Password = "xxx"
	mailConf.SMTPAddr = `smtp.xx.com`
	mailConf.SMTPPort = 25

	m := gomail.NewMessage()
	m.SetHeader(`From`, mailConf.Sender)
	m.SetHeader(`To`, mailConf.RecipientList...)
	m.SetHeader(`Subject`, mailConf.Title)
	m.SetBody(`text/html`, mailConf.Body)
	//m.Attach("./Dockerfile") //添加附件
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	dialer := gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := dialer.DialAndSend(m)
	if err != nil {
		log.Fatalf("Send Email Fail, %s", err.Error())
		return
	}
	log.Printf("Send Email Success")
}