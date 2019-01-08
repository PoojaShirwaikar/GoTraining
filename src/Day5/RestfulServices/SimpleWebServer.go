package main

import (
	"io"
	"net/http"
)

func WriteData(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "writing response to server..")
}

func main() {
	//handle routing
	//if ur is http://localhost:8080/
	http.HandleFunc("/", WriteData)

	//listen to some port
	http.ListenAndServe(":8080", nil)
}
