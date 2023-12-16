package chat

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Room struct {
	roomName   string
	users      map[*User]bool
	usersMutex sync.Mutex
	messageCh  chan Message
}

func NewRoom(roomName string) *Room {
	return &Room{
		roomName:   roomName,
		users:      make(map[*User]bool),
		usersMutex: sync.Mutex{},
		messageCh:  make(chan Message),
	}
}

func (r *Room) HandleMessages() {
	for message := range r.messageCh {
		r.broadcastMessage(&message)
	}
}

func (r *Room) broadcastMessage(message *Message) {
	jsonMessage, _ := json.Marshal(message.Content)
	for user := range r.users {
		err := user.Conn.WriteMessage(websocket.TextMessage, jsonMessage)
		if err != nil {
			log.Fatal(err)
		}
	}
}
