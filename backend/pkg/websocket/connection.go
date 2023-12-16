package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/rajenderK7/go-chat/backend/internal/chat"
)

func HandleConnections(w http.ResponseWriter, r *http.Request, upgrader *websocket.Upgrader, cs *chat.ChatService, roomName string, username string) {
	// Upgrade the HTTP connection to WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	user := chat.NewUser(username, conn)
	cs.JoinRoom(roomName, user)

	// Listen to incoming messages in the conn (i.e. form the user).
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			// If there is an error while reading from a connection
			// it might be disconnected.
			// We want to remove the connection from the room.
			cs.LeaveRoom(roomName, user)
			return
		}

		cs.SendMessage(roomName, messageType, p, user)
	}
}
