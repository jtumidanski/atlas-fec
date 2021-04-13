package consumers

import (
	"atlas-fec/expression"
	"log"
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
	return func(l *log.Logger, e interface{}) {
		if event, ok := e.(*mapChangedEvent); ok {
			expression.GetCache().Clear(event.CharacterId)
		} else {
			l.Printf("[ERROR] unable to cast event provided to handler [MapChangedEvent]")
		}
	}
}
