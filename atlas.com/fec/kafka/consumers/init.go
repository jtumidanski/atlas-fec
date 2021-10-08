package consumers

import (
	"atlas-fec/kafka/handler"
	"context"
	"github.com/sirupsen/logrus"
	"sync"
)

const (
	ChangeFacialExpressionCommand = "change_facial_expression_command"
	ChangeMapEvent                = "change_map_event"
)

func CreateEventConsumers(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup) {
	cec := func(topicToken string, name string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
		createEventConsumer(l, ctx, wg, name, topicToken, emptyEventCreator, processor)
	}
	cec("CHANGE_FACIAL_EXPRESSION", ChangeFacialExpressionCommand, CharacterExpressionCreator(), HandleCharacterExpression())
	cec("TOPIC_CHANGE_MAP_EVENT", ChangeMapEvent, ChangeMapEventCreator(), HandleChangeMapEvent())
}

func createEventConsumer(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup, name string, topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
	wg.Add(1)
	go NewConsumer(l, ctx, wg, name, topicToken, "Facial Expression Service", emptyEventCreator, processor)
}
