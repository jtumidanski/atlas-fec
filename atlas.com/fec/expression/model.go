package expression

import "time"

type Model struct {
	characterId uint32
	mapId       uint32
	expression  uint32
	expiration  time.Time
}

func (m Model) Expiration() time.Time {
	return m.expiration
}

func (m Model) CharacterId() uint32 {
	return m.characterId
}

func (m Model) MapId() uint32 {
	return m.mapId
}

func (m Model) Expression() uint32 {
	return m.expression
}
