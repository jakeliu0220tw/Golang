package main

import (
	"context"
	"net/http"

	"go.uber.org/fx"
)

func register(lifecycle fx.Lifecycle) {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go server.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return server.Shutdown(ctx)
			},
		},
	)
}

func demo2() {
	// Fx has lifecycle for its application
	// Lifecycle allows you to register functions that will be executed at application's start and stop time.
	fx.New(
		// register() is invoked by the app using fx.Invoke()
		fx.Invoke(register),
	).Run()
}
