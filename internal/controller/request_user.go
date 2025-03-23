package controller

import (
	"cuti/internal/model/web"
	"cuti/internal/service"

	"github.com/gofiber/fiber/v2"
)

func RequestCreateUser(c *fiber.Ctx) error {

	userRequest := web.UserRequest{}
	if err := c.BodyParser(&userRequest); err != nil {
		// errStatus := c.Status(fiber.StatusBadRequest).SendString("error format request")

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"kode": "404"})
	}

	service.InputUser(userRequest)

	return c.JSON(fiber.Map{
		"message": "user created Succesfully",
		// "error":   errRes.Error(),
	})

}
