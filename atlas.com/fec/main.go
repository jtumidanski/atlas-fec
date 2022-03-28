package main

import (
	"atlas-fec/character"
	"atlas-fec/character/expression"
	"atlas-fec/kafka"
	"atlas-fec/logger"
	tasks "atlas-fec/task"
	"atlas-fec/tracing"
	"context"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const serviceName = "atlas-fec"
const consumerGroupId = "Facial Expression Service"

func main() {
	l := logger.CreateLogger(serviceName)
	l.Infoln("Starting main service.")

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	tc, err := tracing.InitTracer(l)(serviceName)
	if err != nil {
		l.WithError(err).Fatal("Unable to initialize tracer.")
	}
	defer func(tc io.Closer) {
		err := tc.Close()
		if err != nil {
			l.WithError(err).Errorf("Unable to close tracer.")
		}
	}(tc)

	kafka.CreateConsumers(l, ctx, wg,
		character.ChangeFacialExpressionConsumer(consumerGroupId),
		character.MapChangedConsumer(consumerGroupId))

	go tasks.Register(expression.NewRevertTask(l, time.Millisecond*50))

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
