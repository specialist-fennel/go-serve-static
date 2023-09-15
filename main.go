package main

import (
	"net/http"
)

func main() {

	// Here is your static file server, which serves files from the static/ directory which contains your js and css files
	// usually I put CSS, images and js in here.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	// I've moved the html files into a new directory called templates/

	// Here are your routes for the website with "/" being the home page.
	// I like to use HandleFunc and then make the functions themselves in a separate file. In this case they are in handlers.go
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/chat", ChatPage)
	http.HandleFunc("/example", RenderExample)

	http.ListenAndServe(":3000", nil)
}
