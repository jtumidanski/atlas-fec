package producers

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type characterExpressionChangedEvent struct {
	CharacterId uint32 `json:"characterId"`
	MapId       uint32 `json:"mapId"`
	Expression  uint32 `json:"expression"`
}

func CharacterExpressionChanged(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, mapId uint32, expression uint32) {
	producer := ProduceEvent(l, span, "EXPRESSION_CHANGED")
	return func(characterId uint32, mapId uint32, expression uint32) {
		event := &characterExpressionChangedEvent{CharacterId: characterId, MapId: mapId, Expression: expression}
		producer(CreateKey(int(characterId)), event)
	}
}
