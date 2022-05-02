package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func reqHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	time.Sleep(3 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("Server Timeout")
		w.Write([]byte("<h1>Server Timeout</h1>"))
		return
	default:
		fmt.Println("Hello")
		w.Write([]byte("<h1>Hello</h1>"))
	}
}

func main() {
	http.HandleFunc("/", reqHandler)
	http.ListenAndServe(":8080", nil)
}
