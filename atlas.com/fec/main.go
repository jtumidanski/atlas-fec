package main

import (
	"atlas-fec/kafka/consumers"
	"atlas-fec/logger"
	tasks "atlas-fec/task"
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	l := logger.CreateLogger()
	l.Infoln("Starting main service.")

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	consumers.CreateEventConsumers(l, ctx, wg)

	go tasks.Register(tasks.NewExpressionRevert(l, time.Millisecond*50))

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Infof("Initiating shutdown with signal %s.", sig)
	cancel()
	wg.Wait()
	l.Infoln("Service shutdown.")
}
