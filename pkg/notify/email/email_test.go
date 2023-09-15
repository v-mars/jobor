package email

import (
	"testing"
)

func TestNewSMTP(t *testing.T) {
	// odftxhxilvfzcbci

	mail := NewMail("xx@qq.com", "xx", "smtp.qq.com",
		"xxx", 465, false)
	_ = mail.Send("测试用gomail发送邮件", "Good Good Study, Day Day Up!!!!!!", []string{`xx@qq.com`, "xxx"}, []string{})
}
