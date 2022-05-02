package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := []byte(`{"message":"hello"}`)
	w.Write(msg)
}

func (srv server) Shutdown(ctx context.Context) error {
	fmt.Println("close workers")
	fmt.Println("close db connection")
	return nil
}

func main() {
	// launch HttpServer with goroutine
	srv := server{}
	go func() {
		if err := http.ListenAndServe(":8080", srv); err != nil {
			fmt.Println(err)
		}
	}()

	// make signal channel to receive os quit signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("graceful shutdown ...")

	// close all ctx & resource
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx.Done()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println(err)
	}
}
