package main

import (
	"log"
	"logistica/app/middlewares"
	"logistica/app/routes"
	"logistica/app/utility"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/mustache/v2"
)

var store = *session.New()

func main() {
	utility.ClearScreen()

	engine := mustache.New("./views", ".mustache")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")
	// Middleware global untuk menonaktifkan caching
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("noCache") == "true"
		},
		Expiration:   0 * time.Minute,
		CacheControl: true,
	}))

	// Routes
	routes.SetupRoutes(app, &store)
	app.Get("/metrics", monitor.New())

	// Middlewares
	middlewares.UserAuthorization(app, &store) // Menangani autorisasi user

	// Menjalankan server pada port 3000
	log.Fatal(app.Listen(":1500"))
}
