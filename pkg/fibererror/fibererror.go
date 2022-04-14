package fibererror

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Message string `json:"message"`
}

func New(c *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}
	bb, err := json.Marshal(Message{
		Message: err.Error(),
	})
	if err != nil {
		return err
	}
	_, err = c.Write(bb)
	return err
}