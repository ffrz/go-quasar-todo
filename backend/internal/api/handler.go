package api

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/you/golang-quasar-todo-starter/backend/internal/todo"
)

func RegisterRoutes(app *fiber.App, db *sql.DB) {
	v := app.Group("/api/v1")

	// migrate on start (simple)
	_ = todo.Migrate(db)

	v.Get("/todos", func(c *fiber.Ctx) error {
		list, err := todo.List(db)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(list)
	})

	v.Post("/todos", func(c *fiber.Ctx) error {
		var body struct {
			Title string `json:"title"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		id, err := todo.Create(db, body.Title)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"id": id})
	})

	v.Put("/todos/:id/toggle", func(c *fiber.Ctx) error {
		idS := c.Params("id")
		id, _ := strconv.ParseInt(idS, 10, 64)
		var body struct {
			Done bool `json:"done"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		if err := todo.Toggle(db, id, body.Done); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.SendStatus(204)
	})

	v.Delete("/todos/:id", func(c *fiber.Ctx) error {
		idS := c.Params("id")
		id, _ := strconv.ParseInt(idS, 10, 64)
		if err := todo.Delete(db, id); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.SendStatus(204)
	})

	// optional: static files serve
	app.Static("/", "../frontend/dist")
}
