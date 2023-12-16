package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	ws "github.com/gorilla/websocket"
	"github.com/rajenderK7/go-chat/backend/internal/chat"
	"github.com/rajenderK7/go-chat/backend/pkg/websocket"
)

func main() {
	cs := chat.NewChatService()
	upgrader := ws.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// Allow all connections for simplicity.
			// You might want to implement a more secure CheckOrigin function in production.
			return true
		},
	}

	router := mux.NewRouter()
	// Handle WebSocket connections for a specific room and username
	router.HandleFunc("/ws/{roomName}/{username}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		roomName := vars["roomName"]
		username := vars["username"]

		websocket.HandleConnections(w, r, &upgrader, cs, roomName, username)
	})

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	// Graceful shutdown
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-stop

		log.Println("Shutting down gracefully...")
		if err := server.Shutdown(context.TODO()); err != nil {
			log.Printf("Error during shutdown: %v", err)
		}
	}()

	log.Printf("Server is running on %s...\n", "3000")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %v", err)
	}
}
