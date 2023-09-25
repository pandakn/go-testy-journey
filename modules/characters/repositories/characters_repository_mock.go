package repositories

import (
	"github.com/pandakn/go-testy-journey/modules/characters"
	"github.com/stretchr/testify/mock"
)

type charactersRepositoryMock struct {
	mock.Mock
}

func NewCharacterRepositoryMock() *charactersRepositoryMock {
	return &charactersRepositoryMock{}
}

func (m *charactersRepositoryMock) GetCharacterByName(name string) (*characters.Character, error) {
	args := m.Called(name)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*characters.Character), args.Error(1)
}
