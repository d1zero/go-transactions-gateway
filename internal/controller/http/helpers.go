package http

import "github.com/gofiber/fiber/v2"

func NewResponse(data, metadata any) fiber.Map {
	return fiber.Map{
		"metadata": metadata,
		"data":     data,
	}
}

func NewErrResponse(err error) fiber.Map {
	return fiber.Map{
		"message": err.Error(),
	}
}
