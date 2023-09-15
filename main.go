package main

import (
	"net/http"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir("static")))
	// http.Handle("/", http.StripPrefix(".html", http.FileServer(http.Dir("static"))))
	// http.Handle("/chat", http.FileServer(http.Dir("static/chat.html")))

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.FileServer(http.Dir("static")).ServeHTTP(w, r)
	// })

	// http.HandleFunc("/chat/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.FileServer(http.Dir("./static/chat.hmtl")).ServeHTTP(w, r)
	// })

	// http.Handle("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
	// })

	http.ListenAndServe(":3000", nil)
}
