package wsMessage

import (
	"dumbdumbChat/chatAI"
	"dumbdumbChat/model"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	HandshakeTimeout:  10 * time.Second,
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func UpgradeConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Upgrade wubsocket")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}

	connectedWS = append(connectedWS, conn)
	go receiveMessage(conn)
	go loadChatHistory()
}

func loadChatHistory() {
	chatHistory := chatAI.GetChatHistory()

	for _, chat := range chatHistory {
		if chat.Role != "system" {
			SendMessage(model.Chat, chat)
		}
	}
}
