package consumers

import (
	"atlas-fec/kafka/producers"
	"atlas-fec/rest/requests"
	"context"
	"log"
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
	return func(l *log.Logger, e interface{}) {
		if event, ok := e.(characterExpressionEvent); ok {
			character, err := requests.Character().GetCharacterAttributesById(event.CharacterId)
			if err != nil {
				l.Printf("[ERROR] %s", err.Error())
				return
			}

			producers.CharacterExpressionChanged(l, context.Background()).Emit(event.CharacterId, character.Data().Attributes.MapId, event.Emote)


		} else {
			l.Printf("[ERROR] unable to cast event provided to handler [CharacterExpressionEvent]")
		}
	}
}