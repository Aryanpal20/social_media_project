package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients []websocket.Conn

func main() {

	// create  endpoint for websocket
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// initlize config
		conn, _ := upgrader.Upgrade(w, r, nil)

		clients = append(clients, *conn)

		// loop if client send to server
		for {
			// read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			// prient msg in your console terminal
			fmt.Printf("%s send: %s\n", conn.RemoteAddr(), string(msg))

			// loop if msg found and send again to client for write in your browser
			for _, client := range clients {
				if err = client.WriteMessage(msgType, msg); err != nil {
					return
				}
			}
		}
	})
	// send you html file for open to browser
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
		// w, r is write and delete your index.html
	})
	fmt.Println("your server run 8080")
	http.ListenAndServe(":8080", nil)

}
