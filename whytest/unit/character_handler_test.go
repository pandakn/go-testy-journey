package unit

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/pandakn/go-testy-journey/modules/characters"
	"github.com/pandakn/go-testy-journey/modules/characters/handlers"
	"github.com/pandakn/go-testy-journey/modules/characters/usecases"
	"github.com/pandakn/go-testy-journey/whytest"
	"github.com/stretchr/testify/assert"
)

func TestHandlerGetCharacterByName(t *testing.T) {
	type testCase struct {
		name               string
		characterName      string
		statusCodeExpected int
		isErr              bool
		expected           string
	}

	cases := []testCase{
		{
			name:               "success 200-Ok",
			characterName:      "Panda",
			isErr:              false,
			statusCodeExpected: 200,
			expected: whytest.JsonToString(&characters.Character{
				Id:      "c-0001",
				Name:    "Panda",
				Health:  100,
				Attack:  50,
				Defense: 20,
			}),
		},
		{
			name:               "error 404-not-found",
			characterName:      "gg",
			isErr:              true,
			statusCodeExpected: 404,
			expected:           characters.ErrCharacterNotFound.Error(),
		},
	}

	useCaseMock := usecases.NewCharacterUsecaseMock()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.isErr {
				useCaseMock.On("GetCharacterByName", tc.characterName).Return(nil, characters.ErrCharacterNotFound)
			} else {
				expectedCharacter := &characters.Character{
					Id:      "c-0001",
					Name:    "Panda",
					Health:  100,
					Attack:  50,
					Defense: 20,
				}
				useCaseMock.On("GetCharacterByName", tc.characterName).Return(expectedCharacter, nil)
			}

			handler := handlers.NewCharacterHandler(useCaseMock)

			app := fiber.New()
			app.Get("/character", handler.GetCharacter)

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/character?name=%v", tc.characterName), nil)
			res, _ := app.Test(req)
			defer res.Body.Close()

			//Assert
			if assert.Equal(t, tc.statusCodeExpected, res.StatusCode) {
				body, _ := io.ReadAll(res.Body)
				assert.Equal(t, tc.expected, string(body))
			}
		})
	}
}
