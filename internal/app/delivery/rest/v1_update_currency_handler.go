package rest

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"

	"currency/internal/app/types"
	"currency/pkg/fibererror"
)

type V1UpdateCurrencyRequest struct {
	From string `json:"from"`
	To string `json:"to"`
	Well  float64 `json:"well"`
}

func (h Handler) v1UpdateCurrencyHandler(c *fiber.Ctx) error {
	req := V1UpdateCurrencyRequest{}
	err := json.Unmarshal(c.Request().Body(), &req)
	if err != nil {
		return fibererror.New(c, err)
	}

	err = h.usecase.Pair.Update(c.Context(), types.Currency{
		From: req.From,
		To: req.To,

		Well: req.Well,
	})

	return fibererror.New(c, err)
}
