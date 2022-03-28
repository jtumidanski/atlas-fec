package expression

import (
	"atlas-fec/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type expressionChangedEvent struct {
	CharacterId uint32 `json:"characterId"`
	MapId       uint32 `json:"mapId"`
	Expression  uint32 `json:"expression"`
}

func emitExpressionChanged(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, mapId uint32, expression uint32) {
	producer := kafka.ProduceEvent(l, span, "EXPRESSION_CHANGED")
	return func(characterId uint32, mapId uint32, expression uint32) {
		event := &expressionChangedEvent{CharacterId: characterId, MapId: mapId, Expression: expression}
		producer(kafka.CreateKey(int(characterId)), event)
	}
}
