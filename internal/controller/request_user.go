package controller

import (
	"cuti/internal/model/web"
	"cuti/internal/service"

	"github.com/gofiber/fiber/v2"
)

func RequestCreateUser(c *fiber.Ctx) error {

	userRequest := web.UserRequest{}
	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("ssss")
	}

	service.InputUser(userRequest)

	return c.JSON(fiber.Map{
		"message": "user created Succesfully",
		"user":    userRequest,
	})

}
