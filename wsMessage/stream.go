package wsMessage

import (
	"encoding/json"
	"fmt"
	"live2dViewer/model"
	"log"

	"github.com/gorilla/websocket"
)

var connectedWS = make([]*websocket.Conn, 0)
var receiveMessageGroup = make([]chan string, 0)

func SendMessage(messageType model.WSMessageType, data interface{}) error {
	message := model.WSResponse{
		MessageType: messageType,
		Data:        data,
	}
	wsMessage, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, ws := range connectedWS {
		ws.WriteMessage(websocket.TextMessage, wsMessage)
	}

	return nil
}

func receiveMessage(ws *websocket.Conn) {
	for {
		if ws != nil {
			_, message, err := ws.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("error: %v", err)
				}
				break
			}
			fmt.Println(string(message))

			for _, c := range receiveMessageGroup {
				c <- string(message)
			}
		}
	}

}

func RegisterReceiveStream(c chan string) {
	receiveMessageGroup = append(receiveMessageGroup, c)
}
