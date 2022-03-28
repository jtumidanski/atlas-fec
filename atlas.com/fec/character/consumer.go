package character

import (
	"atlas-fec/character/expression"
	"atlas-fec/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	consumerNameChangeFacialExpression = "change_facial_expression_command"
	consumerNameChangeMap              = "change_map_event"
	topicTokenChangeFacialExpression   = "CHANGE_FACIAL_EXPRESSION"
	topicTokenChangeMap                = "TOPIC_CHANGE_MAP_EVENT"
)

func ChangeFacialExpressionConsumer(groupId string) kafka.ConsumerConfig {
	return kafka.NewConsumerConfig[expressionEvent](consumerNameChangeFacialExpression, topicTokenChangeFacialExpression, groupId, handleCharacterExpression())
}

type expressionEvent struct {
	CharacterId uint32 `json:"characterId"`
	Emote       uint32 `json:"emote"`
}

func handleCharacterExpression() kafka.HandlerFunc[expressionEvent] {
	return func(l logrus.FieldLogger, span opentracing.Span, event expressionEvent) {
		mapId := GetMap(l, span)(event.CharacterId)
		expression.Change(l, span)(event.CharacterId, mapId, event.Emote)
	}
}

func MapChangedConsumer(groupId string) kafka.ConsumerConfig {
	return kafka.NewConsumerConfig[mapChangedEvent](consumerNameChangeMap, topicTokenChangeMap, groupId, handleChangeMap())
}

type mapChangedEvent struct {
	WorldId     byte   `json:"worldId"`
	ChannelId   byte   `json:"channelId"`
	MapId       uint32 `json:"mapId"`
	PortalId    uint32 `json:"portalId"`
	CharacterId uint32 `json:"characterId"`
}

func handleChangeMap() kafka.HandlerFunc[mapChangedEvent] {
	return func(l logrus.FieldLogger, span opentracing.Span, event mapChangedEvent) {
		expression.Clear(event.CharacterId)
	}
}
