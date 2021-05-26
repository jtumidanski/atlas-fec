package requests

import (
	"atlas-fec/rest/attributes"
	"fmt"
)

const (
	charactersServicePrefix string = "/ms/cos/"
	charactersService              = BaseRequest + charactersServicePrefix
	charactersResource             = charactersService + "characters"
	charactersById                 = charactersResource + "/%d"
)

var Character = func() *character {
	return &character{}
}

type character struct {
}

func (c *character) GetCharacterAttributesById(characterId uint32) (*attributes.CharacterAttributesDataContainer, error) {
	ar := &attributes.CharacterAttributesDataContainer{}
	err := get(fmt.Sprintf(charactersById, characterId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
