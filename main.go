package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	var err error

	dsn := "host=postgres user=atiuser password=atipassword dbname=atiweb port=5432 sslmode=disable TimeZone=Asia/Istanbul"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could no connect with the database")

	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World my name is Ahmet👋!")
	})
	app.Listen(":8000")
}
