package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"

	"currency/internal/app/service"
)

type Handler struct {
	usecase *service.UseCase

	fiberAPI *fiber.App
}

func New(usecase *service.UseCase) *Handler {
	h := &Handler{
		usecase: usecase,
	}

	api := fiber.New()
	api.Get("/api/v1/currency",  h.v1GetCurrencyHandler)
	api.Put("/api/v1/currency",  h.v1UpdateCurrencyHandler)
	api.Post("/api/v1/currency", h.v1CreateCurrencyHandler)

	h.fiberAPI = api
	return h
}

func (h *Handler) Listen(port string) error {
	return errors.Wrap(h.fiberAPI.Listen(port), "listen handler api")
}
