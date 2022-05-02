package main

import (
	"log"
	"net/http"
	"time"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get authorization
	auth := r.Header.Get("Authorization")
	if auth != "secret" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Authorization Fail"))
		return
	}

	// pass authorization
	time.Sleep(time.Second * 1)
	w.Write([]byte("hello client"))
}

func main() {
	log.Fatal(http.ListenAndServe(":8999", server{}))
}
