package unit

import (
	"testing"

	"github.com/pandakn/go-testy-journey/modules/characters"
	"github.com/pandakn/go-testy-journey/modules/characters/repositories"
	"github.com/pandakn/go-testy-journey/modules/characters/usecases"
	"github.com/pandakn/go-testy-journey/whytest"
	"github.com/stretchr/testify/assert"
)

func TestUseCaseGetCharacterByName(t *testing.T) {

	type testCase struct {
		name          string
		characterName string
		isErr         bool
		expected      string
	}

	cases := []testCase{
		{
			name:          "success",
			characterName: "Monoserius",
			isErr:         false,
			expected: whytest.JsonToString(&characters.Character{
				Id:      "c-0001",
				Name:    "Monoserius",
				Health:  100,
				Attack:  50,
				Defense: 20,
			}),
		},
		{
			name:          "error not found",
			characterName: "DangGuitar",
			isErr:         true,
			expected:      characters.ErrCharacterNotFound.Error(),
		},
	}
	repo := repositories.NewCharacterRepositoryMock()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Configure the mock repository's behavior based on the test case.
			if tc.isErr {
				repo.On("GetCharacterByName", tc.characterName).Return(nil, characters.ErrCharacterNotFound)
			} else {
				character := &characters.Character{
					Id:      "c-0001",
					Name:    "Monoserius",
					Health:  100,
					Attack:  50,
					Defense: 20,
				}
				repo.On("GetCharacterByName", tc.characterName).Return(character, nil)
			}

			useCase := usecases.NewCharacterUsecase(repo)

			character, err := useCase.GetCharacterByName(tc.characterName)

			if tc.isErr {
				assert.ErrorIs(t, err, characters.ErrCharacterNotFound)
			} else {
				assert.Equal(t, tc.expected, whytest.JsonToString(&character))
			}
		})
	}

}
