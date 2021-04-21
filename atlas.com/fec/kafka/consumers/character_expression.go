package consumers

import (
	"atlas-fec/expression"
	"atlas-fec/kafka/producers"
	"atlas-fec/rest/requests"
	"context"
	"github.com/sirupsen/logrus"
)

type characterExpressionEvent struct {
	CharacterId uint32 `json:"characterId"`
	Emote       uint32 `json:"emote"`
}

func CharacterExpressionCreator() EmptyEventCreator {
	return func() interface{} {
		return &characterExpressionEvent{}
	}
}

func HandleCharacterExpression() EventProcessor {
	return func(l logrus.FieldLogger, e interface{}) {
		if event, ok := e.(*characterExpressionEvent); ok {
			character, err := requests.Character().GetCharacterAttributesById(event.CharacterId)
			if err != nil {
				l.WithError(err).Errorf("Unable to locate the character.")
				return
			}
			mapId := character.Data().Attributes.MapId
			producers.CharacterExpressionChanged(l, context.Background()).Emit(event.CharacterId, mapId, event.Emote)
			expression.GetCache().Add(event.CharacterId, mapId, 0)
		} else {
			l.Errorf("Unable to cast event provided to handler.")
		}
	}
}
