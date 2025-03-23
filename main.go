package main

import (
	"cuti/internal/connection"
	"cuti/internal/controller"
	"cuti/internal/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := connection.OpenConnection()
	errDb := db.AutoMigrate(&model.User{})
	if errDb != nil {
		fmt.Println(errDb.Error())
	}

	app := fiber.New()
	app.Post("/user", controller.RequestCreateUser)

	app.Listen(":3000")
}
