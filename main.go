package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"vayer-electric-backend/env"
	"vayer-electric-backend/gracefulserver"

	"github.com/go-chi/chi/v5"
)

var httpPort = env.PORT

// Returns a context to be the main context of the app.
// It's canceled once a syscall.SIGINT or syscall.SIGTERM is received.
// We could use signal.NotifyContext instead but we won't be able to use the actual signal received for loging & debugging
func getMainContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

		exitSignal := <-signals
		fmt.Printf(fmt.Sprintf("received %s shutting down", exitSignal.String()))

		cancel()
	}()

	return ctx
}

func main() {
	// TODO: add a adecuate logger

	defer func() { fmt.Printf("bye") }()

	mainCtx := getMainContext()

	r := chi.NewRouter()

    server := gracefulserver.New(&http.Server{
		Addr:    fmt.Sprintf(":%d", httpPort),
		Handler: r,
	})

	if err := server.StartListening(mainCtx); err != nil {
		fmt.Printf("server failed to start")
		return
	}

	fmt.Printf("started")

	defer func() {
		fmt.Printf("server stopping")
		server.Shutdown()
		fmt.Printf("server stopped")
	}()

	<-mainCtx.Done()

}