package expression

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func Change(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, mapId uint32, expression uint32) {
	return func(characterId uint32, mapId uint32, expression uint32) {
		emitExpressionChanged(l, span)(characterId, mapId, expression)
		getCache().add(characterId, mapId, 0)
	}
}

func Clear(characterId uint32) {
	getCache().clear(characterId)
}
