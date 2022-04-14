package rest

import (
	"github.com/gofiber/fiber/v2"

	"currency/pkg/fibererror"
)

type V1CurrencyResponse struct {
	Pairs []Pair `json:"pairs"`
}

type Pair struct {
	From string `json:"from"`
	To string `json:"to"`
}

func (h Handler) v1GetCurrencyHandler(c *fiber.Ctx) error {
	currencies, err := h.usecase.Pair.GetAll(c.Context())
	if err != nil {
		return fibererror.New(c, err)
	}

	resp := V1CurrencyResponse{
		Pairs: make([]Pair, 0, len(currencies)),
	}

	for _, cc := range currencies {
		resp.Pairs = append(resp.Pairs, Pair{
			From: cc.From,
			To: cc.To,
		})
	}


	return c.JSON(resp)
}
