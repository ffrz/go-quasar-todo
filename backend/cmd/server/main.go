package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/you/golang-quasar-todo-starter/backend/infra"
	"github.com/you/golang-quasar-todo-starter/backend/internal/api"
)

func main() {
	// load env
	_ = godotenv.Load()

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./data/todos.db"
	}

	db, err := infra.NewSQLite(dbPath)
	if err != nil {
		log.Fatalf("db init: %v", err)
	}
	defer db.Close()

	app := fiber.New()

	root := app.Group("/")
	root.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!!")
	})

	api.RegisterRoutes(app, db)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("listening on :" + port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("listen: %v", err)
	}
}
