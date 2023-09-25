package usecases

import (
	"github.com/pandakn/go-testy-journey/modules/characters"
	"github.com/stretchr/testify/mock"
)

type charactersUsecaseMock struct {
	mock.Mock
}

func NewCharacterUsecaseMock() *charactersUsecaseMock {
	return &charactersUsecaseMock{}
}

func (m *charactersUsecaseMock) GetCharacterByName(name string) (*characters.Character, error) {
	args := m.Called(name)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*characters.Character), args.Error(1)
}
