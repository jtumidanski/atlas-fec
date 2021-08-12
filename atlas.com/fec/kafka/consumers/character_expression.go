package consumers

import (
	"atlas-fec/character"
	"atlas-fec/expression"
	"atlas-fec/kafka/handler"
	"atlas-fec/kafka/producers"
	"github.com/sirupsen/logrus"
)

type characterExpressionEvent struct {
	CharacterId uint32 `json:"characterId"`
	Emote       uint32 `json:"emote"`
}

func CharacterExpressionCreator() handler.EmptyEventCreator {
	return func() interface{} {
		return &characterExpressionEvent{}
	}
}

func HandleCharacterExpression() handler.EventHandler {
	return func(l logrus.FieldLogger, e interface{}) {
		if event, ok := e.(*characterExpressionEvent); ok {
			mapId := character.GetMap(l)(event.CharacterId)
			producers.CharacterExpressionChanged(l)(event.CharacterId, mapId, event.Emote)
			expression.GetCache().Add(event.CharacterId, mapId, 0)
		} else {
			l.Errorf("Unable to cast event provided to handler.")
		}
	}
}
