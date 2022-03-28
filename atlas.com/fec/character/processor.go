package character

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func GetMap(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) uint32 {
	return func(characterId uint32) uint32 {
		c, err := requestById(characterId)(l, span)
		if err != nil {
			l.WithError(err).Errorf("Unable to locate the character.")
			return 0
		}
		attr := c.Data().Attributes
		return attr.MapId
	}
}
