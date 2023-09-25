package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pandakn/go-testy-journey/modules/characters/usecases"
)

type ICharacterHandler interface {
	GetCharacter(c *fiber.Ctx) error
}

type characterHandler struct {
	charactersUsecases usecases.ICharacterUsecase
}

func NewCharacterHandler(charactersUsecases usecases.ICharacterUsecase) ICharacterHandler {
	return &characterHandler{
		charactersUsecases: charactersUsecases,
	}
}

func (h *characterHandler) GetCharacter(c *fiber.Ctx) error {
	name := c.Query("name")

	character, err := h.charactersUsecases.GetCharacterByName(name)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(character)
}
