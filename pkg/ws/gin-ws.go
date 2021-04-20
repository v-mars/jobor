package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var UpGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}



type data struct {
	Msg string `json:"msg"`
}
//webSocket请求ping 返回pong
func WS(c *gin.Context) {
	//升级get请求为webSocket协议
	Ws, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	//Ws, err := UpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer Ws.Close()
	for {
		//fmt.Println("begin ws....")
		//读取ws中的数据
		mt, message, err := Ws.ReadMessage()
		//fmt.Println("ws waiting....")
		if err != nil {
			break
		}
		//msg := &data{}
		//_ = json.Unmarshal(message, msg)
		//fmt.Println("msg:", msg)
		//fmt.Println("mt:", mt)
		//fmt.Println("message:", message)
		//fmt.Println("str message:",string(message))

		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入ws数据
		err = Ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
