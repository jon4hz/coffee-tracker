package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jon4hz/coffee-tracker/internal/database"
)

func NewAPI() *fiber.App {
	app := fiber.New()
	app.Get("/user/:user/project/:project", func(c *fiber.Ctx) error {
		user, err := strconv.Atoi(c.Params("user"))
		if err != nil {
			return c.Status(400).SendString("Invalid user")
		}
		p, err := database.NewProject(c.Params("project"), int64(user))
		if err != nil {
			return c.Status(400).SendString("Invalid project")
		}
		return c.SendString(strconv.Itoa(int(p.GetCoffees())))
	})
	return app
}
