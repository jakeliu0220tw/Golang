package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type msgData struct {
	Message string `json:"message"`
}

type server struct{}

// demo of post response
func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := msgData{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	data.Message = strings.ToUpper(data.Message)
	byteStr, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(byteStr)
}

func (srv server) ShutDown() {
	fmt.Println("graceful shutdown ...")
}

func main() {
	srv := server{}
	go func() {
		err := http.ListenAndServe(":8080", srv)
		if err != nil {
			fmt.Println(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // blocking until receiving TERM

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	ctx.Done()
	srv.ShutDown()
}
