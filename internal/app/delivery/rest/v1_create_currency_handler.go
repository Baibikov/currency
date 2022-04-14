package rest

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"

	"currency/internal/app/types"
	"currency/pkg/fibererror"
)

type V1CreateCurrencyRequest struct {
	From string `json:"from"`
	To string `json:"to"`
}

func (h Handler) v1CreateCurrencyHandler(c *fiber.Ctx) error {
	req := V1CreateCurrencyRequest{}
	err := json.Unmarshal(c.Request().Body(), &req)
	if err != nil {
		return fibererror.New(c, err)
	}

	err = h.usecase.Pair.Create(c.Context(), types.Currency{
		From: req.From,
		To: req.To,
	})

	return fibererror.New(c, err)}
