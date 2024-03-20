package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ErrorObject struct {
	Msg   string
	Field string
}

func BadRequest(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusBadRequest).JSON(data)
}

func InternalServerError(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusInternalServerError).JSON(data)
}
func NotFound(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusNotFound).JSON(data)
}
