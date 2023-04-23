package app

import (
	"os"

	"github.com/YungBenn/go-mysql-test/config"
	"github.com/YungBenn/go-mysql-test/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func SetupApp() error {
	err := config.LoadENV()
	if err != nil {
		return err
	}

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK!")
	})

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	app.Listen(":" + port)

	return nil
}
