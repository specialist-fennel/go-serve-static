package main

import (
	"net/http"
)

// The HomePage and ChatPage handler functions use the regular http.ServeFile function to serve the html files.

func HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func ChatPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/chat.html")
}

// This function is an example of how I'd usually render html to the client. The magic mostly happens in the RenderWithData function in render.go
func RenderExample(w http.ResponseWriter, r *http.Request) {
	RenderWithData(w, r, "templates/example.html", map[string]string{
		"some_key": "Hello from the server!",
	})
}
