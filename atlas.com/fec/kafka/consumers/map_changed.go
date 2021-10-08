package consumers

import (
	"atlas-fec/expression"
	"atlas-fec/kafka/handler"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type mapChangedEvent struct {
	WorldId     byte   `json:"worldId"`
	ChannelId   byte   `json:"channelId"`
	MapId       uint32 `json:"mapId"`
	PortalId    uint32 `json:"portalId"`
	CharacterId uint32 `json:"characterId"`
}

func ChangeMapEventCreator() handler.EmptyEventCreator {
	return func() interface{} {
		return &mapChangedEvent{}
	}
}

func HandleChangeMapEvent() handler.EventHandler {
	return func(l logrus.FieldLogger, span opentracing.Span, e interface{}) {
		if event, ok := e.(*mapChangedEvent); ok {
			expression.GetCache().Clear(event.CharacterId)
		} else {
			l.Errorf("Unable to cast event provided to handler.")
		}
	}
}
