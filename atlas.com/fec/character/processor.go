package character

import "github.com/sirupsen/logrus"

func GetMap(l logrus.FieldLogger) func(characterId uint32) uint32 {
	return func(characterId uint32) uint32 {
		c, err := requestById(l)(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to locate the character.")
			return 0
		}
		return c.Data().Attributes.MapId
	}
}
