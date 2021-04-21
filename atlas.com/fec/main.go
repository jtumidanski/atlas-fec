package main

import (
	"atlas-fec/kafka/consumers"
	"atlas-fec/logger"
	tasks "atlas-fec/task"
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	l := logger.CreateLogger()

	createEventConsumers(l)

	go tasks.Register(tasks.NewExpressionRevert(l, time.Millisecond*50))

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Infoln("Shutting down via signal:", sig)
}

func createEventConsumers(l *logrus.Logger) {
	cec := func(topicToken string, emptyEventCreator consumers.EmptyEventCreator, processor consumers.EventProcessor) {
		createEventConsumer(l, topicToken, emptyEventCreator, processor)
	}
	cec("CHANGE_FACIAL_EXPRESSION", consumers.CharacterExpressionCreator(), consumers.HandleCharacterExpression())
	cec("TOPIC_CHANGE_MAP_EVENT", consumers.ChangeMapEventCreator(), consumers.HandleChangeMapEvent())
}

func createEventConsumer(l *logrus.Logger, topicToken string, emptyEventCreator consumers.EmptyEventCreator, processor consumers.EventProcessor) {
	h := func(logger logrus.FieldLogger, event interface{}) {
		processor(logger, event)
	}

	c := consumers.NewConsumer(l, context.Background(), h,
		consumers.SetGroupId("Facial Expression Service"),
		consumers.SetTopicToken(topicToken),
		consumers.SetEmptyEventCreator(emptyEventCreator))
	go c.Init()
}
