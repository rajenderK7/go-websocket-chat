package chat

import (
	"fmt"
	"sync"
)

type ChatService struct {
	rooms      map[string]*Room
	roomsMutex sync.Mutex
}

func NewChatService() *ChatService {
	return &ChatService{
		rooms:      make(map[string]*Room),
		roomsMutex: sync.Mutex{},
	}
}

func (cs *ChatService) CreateRoom(roomName string) {
	cs.roomsMutex.Lock()
	defer cs.roomsMutex.Unlock()

	_, exists := cs.rooms[roomName]
	if exists {
		fmt.Printf("Room \"%s\" already exists\n", roomName)
		return
	}

	room := NewRoom(roomName)
	cs.rooms[roomName] = room

	// Listen to messages in the room on a seperate goroutine.
	go room.HandleMessages()
}

func (cs *ChatService) JoinRoom(roomName string, user *User) {
	_, exists := cs.rooms[roomName]
	if !exists {
		cs.CreateRoom(roomName)
	}
	cs.roomsMutex.Lock()
	defer cs.roomsMutex.Unlock()

	room := cs.rooms[roomName]
	room.usersMutex.Lock()
	room.users[user] = true
	room.usersMutex.Unlock()
}

func (cs *ChatService) LeaveRoom(roomName string, user *User) {
	cs.roomsMutex.Lock()
	defer cs.roomsMutex.Unlock()

	room := cs.rooms[roomName]
	room.usersMutex.Lock()
	delete(room.users, user)
	room.usersMutex.Unlock()
}

func (cs *ChatService) SendMessage(roomName string, messageType int, data []byte, user *User) {
	cs.roomsMutex.Lock()
	defer cs.roomsMutex.Unlock()

	room, exists := cs.rooms[roomName]
	if !exists {
		fmt.Printf("Room \"%s\" does not exist", roomName)
		return
	}

	room.messageCh <- *NewMessage(data, user)
}
