package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/nandangrover/go-microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create the handlers
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	// wrapping ListenAndServe in gofunc so it's not going to block
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// make a new channel to notify on os interrupt of server (ctrl + C)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// This blocks the code until the channel receives some message
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	// Once message is consumed shut everything down
	// Gracefully shuts down all client requests. Makes server more reliable
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
