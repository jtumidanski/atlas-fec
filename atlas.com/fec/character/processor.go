package character

import (
	"atlas-fec/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func GetMap(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) (uint32, error) {
	return func(characterId uint32) (uint32, error) {
		return requests.Provider[attributes, uint32](l, span)(requestById(characterId), mapIdGetter)()
	}
}

func mapIdGetter(body requests.DataBody[attributes]) (uint32, error) {
	attr := body.Attributes
	return attr.MapId, nil
}
