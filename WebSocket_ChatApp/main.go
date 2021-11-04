package main

import (
	"WebSocket_ChatApp/socket"
	"log"
	"net/http"
)

func main() {
	// create new ServeMux
	r := http.NewServeMux()
	// serve the views folder for access html,js files
	fs := http.FileServer(http.Dir("./views/"))
	r.Handle("/", http.StripPrefix("", fs))
	// run socketReaderCreate function
	r.HandleFunc("/socket", socket.SocketReaderCreate)

	log.Println("Server Running : http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
