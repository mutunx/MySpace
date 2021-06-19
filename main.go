package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "Welcome to my space!\n")
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
