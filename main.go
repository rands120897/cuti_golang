package main

import (
	"cuti/internal/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=absent port=5432"
	db, errDb := gorm.Open(postgres.Open(dsn))
	if errDb != nil {
		fmt.Printf("errorr")
	}

	return db
}

func main() {
	user := model.User{
		ID: 15,

		Name:  "dito",
		Email: "sakkj@gmail.com",

		Password: "sdakdjakjjka",
	}

	dsn := "host=localhost user=postgres password=postgres dbname=absent port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("errorr")
	}
	result := db.Create(&user)
	errs := result.Error
	fmt.Println(errs)

	app := fiber.New()
	app.Post("/user", func(c *fiber.Ctx) error {

		if err := c.BodyParser(&user); err != nil {
			return c.JSON(fiber.Map{
				"message": "invalid data",
				"error":   fiber.ErrBadGateway,
			})
		}

		return c.JSON(fiber.Map{
			"message":  "user Created Successfully",
			"username": user,
		})

	})

	app.Listen(":3000")

}
