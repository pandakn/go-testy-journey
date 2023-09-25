package repositories

import (
	"github.com/pandakn/go-testy-journey/modules/characters"
)

type ICharacterRepository interface {
	GetCharacterByName(name string) (*characters.Character, error)
}

type charactersRepository struct {
	characters map[string]*characters.Character
}

func NewCharacterRepository() ICharacterRepository {
	return &charactersRepository{
		characters: map[string]*characters.Character{
			"Monoserius": {
				Id:      "c-0001",
				Name:    "Monoserius",
				Health:  100,
				Attack:  50,
				Defense: 20,
			},
			"Grapi": {
				Id:      "c-0002",
				Name:    "Grapi",
				Health:  80,
				Attack:  20,
				Defense: 30,
			},
		},
	}
}

func (r *charactersRepository) GetCharacterByName(name string) (*characters.Character, error) {
	// Simulate retrieving a character from the database.
	character, ok := r.characters[name]
	if !ok {
		return nil, characters.ErrCharacterNotFound
	}
	return character, nil
}
