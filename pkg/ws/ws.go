package ws

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// ClientManager is a websocket manager
type ClientManager struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

// Client is a websocket client
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
}

// Message is an object for websocket message which is mapped to json type
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

// Manager define a ws server manager
var Manager = ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
}

// Start is to start a ws server
func (manager *ClientManager) Start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			manager.Send(jsonMessage, conn)
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.Send(jsonMessage, conn)
			}
		case message := <-manager.Broadcast:
			for conn := range manager.Clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(manager.Clients, conn)
				}
			}
		}
	}
}

// Send is to send ws message to ws client
func (manager *ClientManager) Send(message []byte, ignore *Client) {
	for conn := range manager.Clients {
		if conn != ignore {
			conn.Send <- message
		}
	}
}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			Manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.ID, Content: string(message)})
		Manager.Broadcast <- jsonMessage
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			_ = c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}


//Start():启动websocket服务
//Send():向连接websocket的管道chan写入数据
//Read():读取在websocket管道中的数据
//Write():通过websocket协议向连接到ws的客户端发送数据


//r.GET("/ws", WsPage)
// WsPage is a websocket handler
func WsPage(c *gin.Context) {
	// change the reqest to websocket model
	//conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	//if err != nil {
	//	http.NotFound(c.Writer, c.Request)
	//	return
	//}

	// websocket connect
	//client := &ws.Client{Id: uuid.NewV4().String(), Socket: conn, Send: make(chan []byte)}
	//
	//ws.Manager.Register <- client
	//
	//go client.Read()
	//go client.Write()
}
