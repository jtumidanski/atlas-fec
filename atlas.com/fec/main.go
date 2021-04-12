package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	l := log.New(os.Stdout, "fec ", log.LstdFlags|log.Lmicroseconds)

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Println("[INFO] shutting down via signal:", sig)
}
