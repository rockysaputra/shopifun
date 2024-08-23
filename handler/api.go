package handler

import "github.com/gofiber/fiber/v3"

type Response struct {
	Status  int
	Message string
}

func Hello(c fiber.Ctx) error {

	response := Response{
		Status:  200,
		Message: "Hello Im Ok",
	}

	return c.Status(fiber.StatusAccepted).JSON(response)
}
