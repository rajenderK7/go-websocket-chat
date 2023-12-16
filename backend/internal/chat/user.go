package chat

import "github.com/gorilla/websocket"

type User struct {
	Username string
	Conn     *websocket.Conn
}

func NewUser(username string, conn *websocket.Conn) *User {
	return &User{
		Username: username,
		Conn:     conn,
	}
}
