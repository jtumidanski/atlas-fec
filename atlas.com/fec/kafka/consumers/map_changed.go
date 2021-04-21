package consumers

import (
	"atlas-fec/expression"
	"github.com/sirupsen/logrus"
)

type mapChangedEvent struct {
	WorldId     byte   `json:"worldId"`
	ChannelId   byte   `json:"channelId"`
	MapId       uint32 `json:"mapId"`
	PortalId    uint32 `json:"portalId"`
	CharacterId uint32 `json:"characterId"`
}

func ChangeMapEventCreator() EmptyEventCreator {
	return func() interface{} {
		return &mapChangedEvent{}
	}
}

func HandleChangeMapEvent() EventProcessor {
	return func(l logrus.FieldLogger, e interface{}) {
		if event, ok := e.(*mapChangedEvent); ok {
			expression.GetCache().Clear(event.CharacterId)
		} else {
			l.Errorf("Unable to cast event provided to handler.")
		}
	}
}
