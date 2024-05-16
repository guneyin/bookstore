package middleware

import "github.com/gofiber/fiber/v2"

type status string

const (
	statusSuccess  status = "SUCCESS"
	statusError    status = "ERROR"
	statusNotfound status = "NOT-FOUND"
)

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func OK(c *fiber.Ctx, msg string, data any) error {
	return c.Status(fiber.StatusOK).JSON(response{
		Status:  string(statusSuccess),
		Message: msg,
		Data:    data,
	})
}

func Error(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(response{
		Status:  string(statusSuccess),
		Message: err.Error(),
		Data:    nil,
	})
}

func NotFound(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusNotFound).JSON(response{
		Status:  string(statusSuccess),
		Message: err.Error(),
		Data:    nil,
	})
}
