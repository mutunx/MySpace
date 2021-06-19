package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "Welcome to my space!\n")
	})
	http.ListenAndServe(":80", nil)
}
