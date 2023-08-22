package email

import (
	"crypto/tls"
	"fmt"
	"jobor/pkg/logger"
	"log"

	"gopkg.in/gomail.v2"
)

// MailboxConf 邮箱配置
type MailboxConf struct {
	// 邮件标题
	Title string
	// 邮件内容
	Body string
	// 邮件附件
	Attach []string
	// 收件人列表
	RecipientList []string
	// 发件人账号
	Sender string
	// 发件人密码，QQ邮箱这里配置授权码
	Password string
	From     string
	// SMTP 服务器地址， QQ邮箱是smtp.qq.com
	SMTPAddr string
	// SMTP端口 QQ邮箱是25
	SMTPPort int
	Tls      bool
}

func NewMail(username, password, SMTPAddr, from string, SMTPPort int, tls bool) MailboxConf {
	return MailboxConf{Sender: username, Password: password, SMTPAddr: SMTPAddr, SMTPPort: SMTPPort, From: from, Tls: tls}
}

func (mailConf MailboxConf) Send(title, body string, RecipientList, AttachList []string) error {
	m := gomail.NewMessage()
	m.SetHeader(`From`, mailConf.Sender)
	m.SetHeader(`To`, RecipientList...)
	m.SetHeader(`Subject`, title)
	m.SetBody(`text/html`, body)
	//m.Attach("./Dockerfile") //添加附件
	if len(AttachList) > 0 {
		for _, v := range AttachList {
			m.Attach(v)
		}
	}

	if len(RecipientList) == 0 {
		err := fmt.Errorf("Send Email Fail, 邮箱接收者不能为空\n")
		logger.Error(err)
		return err
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

/*

-----------------------------------
120 Service ready in nnn minutes.
服务在NNN时间内可用
-----------------------------------
125 Data connection already open; transfer starting.
数据连接已经打开，开始传送数据.
-----------------------------------
150 File status okay; about to open data connection.
文件状态正确，正在打开数据连接.
-----------------------------------
200 Command okay.
命令执行正常结束.
-----------------------------------
202 Command not implemented, superfluous at this site.
命令未被执行，此站点不支持此命令.
-----------------------------------
211 System status, or system help reply.
系统状态或系统帮助信息回应.
-----------------------------------
212 Directory status.
目录状态信息.
-----------------------------------
213 File status.
文件状态信息.
-----------------------------------
214 Help message.On how to use the server or the meaning of a particular non-standard command. This reply is useful only to the human user. 帮助信息。关于如何使用本服务器或特殊的非标准命令。此回复只对人有用。
-----------------------------------
215 NAME system type. Where NAME is an official system name from the list in the Assigned Numbers document.
NAME系统类型。
-----------------------------------
220 Service ready for new user.
新连接的用户的服务已就绪
-----------------------------------
221 Service closing control connection.
控制连接关闭
-----------------------------------
225 Data connection open; no transfer in progress.
数据连接已打开，没有进行中的数据传送
-----------------------------------
226 Closing data connection. Requested file action successful (for example, file transfer or file abort).
正在关闭数据连接。请求文件动作成功结束（例如，文件传送或终止）
-----------------------------------
227 Entering Passive Mode (h1,h2,h3,h4,p1,p2).
进入被动模式
-----------------------------------
230 User logged in, proceed. Logged out if appropriate.
用户已登入。 如果不需要可以登出。
-----------------------------------
250 Requested file action okay, completed.
被请求文件操作成功完成
-----------------------------------
257 "PATHNAME" created.
路径已建立
-----------------------------------
331 User name okay, need password.
用户名存在，需要输入密码
-----------------------------------
332 Need account for login.
需要登陆的账户
-----------------------------------
350 Requested file action pending further information
对被请求文件的操作需要进一步更多的信息
-----------------------------------
421 Service not available, closing control connection.This may be a reply to any command if the service knows it must shut down.
服务不可用，控制连接关闭。这可能是对任何命令的回应，如果服务认为它必须关闭
-----------------------------------
425 Can’t open data connection.
打开数据连接失败
-----------------------------------
426 Connection closed; transfer aborted.
连接关闭，传送中止。
-----------------------------------
450 Requested file action not taken.
对被请求文件的操作未被执行
-----------------------------------
451 Requested action aborted. Local error in processing.
请求的操作中止。处理中发生本地错误。
-----------------------------------
452 Requested action not taken. Insufficient storage space in system.File unavailable (e.g., file busy).
请求的操作没有被执行。 系统存储空间不足。 文件不可用
-----------------------------------
500 Syntax error, command unrecognized. This may include errors such as command line too long.
语法错误，不可识别的命令。 这可能是命令行过长。
-----------------------------------
501 Syntax error in parameters or arguments.
参数错误导致的语法错误
-----------------------------------
502 Command not implemented.
命令未被执行
-----------------------------------
503 Bad sequence of commands.
命令的次序错误。
-----------------------------------
504 Command not implemented for that parameter.
由于参数错误，命令未被执行
-----------------------------------
530 Not logged in.
没有登录
-----------------------------------
532 Need account for storing files.
存储文件需要账户信息
-----------------------------------
550 Requested action not taken. File unavailable (e.g., file not found, no access).
请求操作未被执行，文件不可用。
-----------------------------------
551 Requested action aborted. Page type unknown.
请求操作中止，页面类型未知
-----------------------------------
552 Requested file action aborted. Exceeded storage allocation (for current directory or dataset).
对请求文件的操作中止。 超出存储分配
-----------------------------------
553 Requested action not taken. File name not allowed
请求操作未被执行。 文件名不允许
-----------------------------------
-----------------------------------
这种错误跟http协议类似，大致是：
2开头－－成功
3开头－－权限问题
4开头－－文件问题
5开头－－服务器问题
对偶们最有用的：
421：一般出现在连接数多，需稍后在接；
530：密码错误；
550：目录或文件已经被删除。
*/
