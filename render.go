package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

// RenderWithData renders a template with data.
func RenderWithData(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	// Here we declare a bytes.Buffer to store the output of the template execution.
	buf := new(bytes.Buffer)

	tmpl, err := template.ParseFiles(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// You can see that the first argument to Execute takes an io.Writer. We could simply pass w directly to Execute, but if we did this, the template would be rendered directly to
	// the ResponseWriter and any error would be written to the ResponseWriter as well. Instead we write to the buffer first and then check if there was an error.
	// If there is, we reply with a 500 Internal Server Error and return from the function.
	err = tmpl.Execute(buf, data)
	// Note the data argument to Execute. You can pass data into the template to be rendered. Check out templates/example.html to see how to access this data.
	// The syntax is a little weird for templates but there are lots of good resources to help you learn.
	// I'd start here https://gowebexamples.com/templates/
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Here we can use the buffer's, that we wrote the template, WriteTo method to write the contents of the buffer to the ResponseWriter.
	// We then check for an error.
	n, err := buf.WriteTo(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Written %d bytes", n)
}
