package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type Message struct {
	Event   string `json:"event"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func NewMessage(event, name, content string) *Message {
	return &Message{
		Event:   event,
		Name:    name,
		Content: content,
	}
}

// 由於透過 WebSocket 傳送訊息要使用 []byte 格式，
// 因此這邊我們也將轉換的方法進行封裝
func (m *Message) GetByteMessage() []byte {
	result, _ := json.Marshal(m)
	return result
}

func main() {
	server := gin.Default()
	m := melody.New()
	server.LoadHTMLGlob("template/html/*")
	server.Static("/assets", "./template/assets")
	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	server.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
	// 處理連線
	// melody 有提供 HandleConnect 方法讓我們處理連線的
	// session，我們這邊設定連線進來就發送一個 xxx 加入聊天室 的訊息給全部人
	m.HandleConnect(func(session *melody.Session) {
		id := session.Request.URL.Query().Get("id")
		m.Broadcast(NewMessage("other", id, "加入聊天室").GetByteMessage())
	})
	// 處理離線
	// melody 有提供 HandleConnect 方法讓我們處理連線的 session，
	// 我們這邊設定連線進來就發送一個 xxx 加入聊天室 的訊息給全部人
	m.HandleClose(func(session *melody.Session, i int, s string) error {
		id := session.Request.URL.Query().Get("id")
		m.Broadcast(NewMessage("other", id, "離開聊天室").GetByteMessage())
		return nil
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	server.Run(":8888")
}
