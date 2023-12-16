# go-websocket-chat

This is a fully functional Go backend for chat applications using WebSockets. 

Primary packages used: Gorilla MUX and Gorilla WebSocket

The modular backend enables swift and effortless customization to meet specific project requirements.

## Run the backend

### Install packages
`go get github.com/gorilla/mux`

`go get github.com/gorilla/websocket`

### Start the server
`cd backend/cmd/chat-server`

`go run main.go`

## Access the server
Go to:
`ws://localhost:PORT/ws/{roomName}/{username}`

## Plug in your favorite frontend
Open a WebSocket connection to the above URL with `roomName` and `username`

Have a nice chat 👋
