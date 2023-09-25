package usecases

import (
	"github.com/pandakn/go-testy-journey/modules/characters"
	"github.com/pandakn/go-testy-journey/modules/characters/repositories"
)

type ICharacterUsecase interface {
	GetCharacterByName(name string) (*characters.Character, error)
}

type characterUsecase struct {
	characterRepository repositories.ICharacterRepository
}

func NewCharacterUsecase(characterRepository repositories.ICharacterRepository) ICharacterUsecase {
	return &characterUsecase{
		characterRepository: characterRepository,
	}
}

func (u *characterUsecase) GetCharacterByName(name string) (*characters.Character, error) {
	character, error := u.characterRepository.GetCharacterByName(name)
	if error != nil {
		return nil, error
	}

	return character, nil
}
