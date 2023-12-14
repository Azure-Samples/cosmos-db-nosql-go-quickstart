package main

import (
	"log"

	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

func main() {
	server := socketio.NewServer(
		&engineio.Options{
			Transports: []transport.Transport{
				&websocket.Transport{
					CheckOrigin: func(r *http.Request) bool {
						return true
					},
				},
			},
		},
	)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("Connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "start", func(s socketio.Conn, msg string) {
		s.SetContext(msg)
		log.Println("Greeting:", msg)
		startCosmos(func (msg string) {
			s.Emit("new_message", msg)
		})
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("Error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("Disconnected:", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Listening at localhost:8000...")
    log.Fatal(http.ListenAndServe(":8000", nil))
}