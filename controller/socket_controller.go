package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	socketio "github.com/googollee/go-socket.io"
)

type SocketController struct{}

type Message struct {
	Type      string `json:"type"`
	ChatType  string `json:"chatType"`
	SendId    string `json:"sendId"`
	ReceiveId string `json:"receiveId"`
	Content   string `json:"content"`
}

func (sc *SocketController) RegisterSocket(engine *gin.Engine, server *socketio.Server) {
	engine.GET("/socket.io/*any", gin.WrapH(server))
	engine.POST("/socket.io/*any", gin.WrapH(server))

	server.OnConnect("/", sc.onConnect)
	server.OnError("/", sc.onError)
	server.OnDisconnect("/", sc.onDisconnect)

	server.OnEvent("/", "SINGLE_CHAT", sc.onSingleChat)
	server.OnEvent("/", "BYE", sc.onBye)
}

func (sc *SocketController) onConnect(s socketio.Conn) error {
	s.SetContext("")
	s.Emit("CONNECT", "欢迎连接 ~ ")
	return nil
}

func (sc *SocketController) onError(s socketio.Conn, e error) {
	fmt.Println("meet error:", e)
}

func (sc *SocketController) onDisconnect(s socketio.Conn, reason string) {
	fmt.Println("closed", reason)
}

// func (sc *SocketController) OnSingleChat(s socketio.Conn, message map[string]string) {
func (sc *SocketController) onSingleChat(s socketio.Conn, message Message) {
	fmt.Printf("%v", message)
	// s.SetContext(msg)
	s.Emit("SINGLE_CHAT", "收到的消息是"+message.Content)
}

func (sc *SocketController) onBye(s socketio.Conn) string {
	last := s.Context().(string)
	s.Emit("BYE", last)
	s.Close()
	return last
}
