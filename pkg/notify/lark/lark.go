package lark

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"jobor/pkg/notify"
	"net/http"
	"time"
)

func GenSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}

type Secure uint

const (
	// CustomKey Custom keywords
	CustomKey Secure = iota + 1
	// Sign need sign up
	Sign
	// IPCdir IP addres
	IPCdir
)

// Lark alarm conf
type Lark struct {
	MsgType    string `json:"msg_type"`
	WebHookUrl string `json:"web_hook_url"`
	Sl         Secure `json:"sl"`
	Secret     string `json:"secret"`
}


// Result post resp
type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Content struct {
	Text string `json:"text"`
}

// SendMsg post json data
type SendMsg struct {
	Timestamp string  `json:"timestamp"`
	Sign      string  `json:"sign"`
	MsgType   string  `json:"msg_type"`
	Content   Content `json:"content"`
}

// NewLark init a Lark send conf
func NewLark(webHookUrl string, sl Secure, secret string) notify.Sender {
	d := Lark{
		WebHookUrl: webHookUrl,
		MsgType:    "text",
		Sl:         sl,
		Secret:     secret,
	}

	return &d
}

// Send to notify tos is phone number
func (d *Lark) Send(tos []string, title string, content string) error {
	var reqUrl = d.WebHookUrl
	timestamp := time.Now().Unix()
	sign := ""
	if len(d.Secret)>0{
		var err error
		if sign, err = GenSign(d.Secret, timestamp);err!=nil{
			return err
		}
	}

	sendMsg := SendMsg{
		Timestamp: fmt.Sprintf("%d",timestamp),
		Sign: sign,
		MsgType: "text",
		Content: Content{
			Text: title + "\n" + content + "\n",
		},
	}

	resp, err := notify.JSONPost(http.MethodPost, reqUrl, sendMsg, http.DefaultClient)
	if err != nil {
		return err
	}
	res := Result{}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return err
	}
	if res.Code != 0 {
		return fmt.Errorf("errmsg: %s errcode: %d", res.Msg, res.Code)
	}
	return nil
}