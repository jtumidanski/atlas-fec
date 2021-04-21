package producers

import (
	"context"
	"github.com/sirupsen/logrus"
)

type characterExpressionChangedEvent struct {
	CharacterId uint32 `json:"characterId"`
	MapId       uint32 `json:"mapId"`
	Expression  uint32 `json:"expression"`
}

var CharacterExpressionChanged = func(l logrus.FieldLogger, ctx context.Context) *characterExpressionChanged {
	return &characterExpressionChanged{
		l, ctx,
	}
}

type characterExpressionChanged struct {
	l   logrus.FieldLogger
	ctx context.Context
}

func (c *characterExpressionChanged) Emit(characterId uint32, mapId uint32, expression uint32) {
	event := &characterExpressionChangedEvent{CharacterId: characterId, MapId: mapId, Expression: expression}
	produceEvent(c.l, "EXPRESSION_CHANGED", createKey(int(characterId)), event)
}
