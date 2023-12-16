package chat

import "time"

type Message struct {
	Sender    *User     `json:"sender"`  // Sender of the message
	Content   string    `json:"content"` // Content of the message
	Timestamp time.Time `json:"timestamp"`
}

func NewMessage(data []byte, user *User) *Message {
	return &Message{
		Sender:    user,
		Content:   string(data),
		Timestamp: time.Now(),
	}
}
