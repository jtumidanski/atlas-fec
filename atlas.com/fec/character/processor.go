package character

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func GetMap(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) uint32 {
	return func(characterId uint32) uint32 {
		c, err := requestById(l, span)(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to locate the character.")
			return 0
		}
		return c.Data().Attributes.MapId
	}
}
