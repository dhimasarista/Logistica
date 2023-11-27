package main

import (
	"log"
	"logistica/app/middlewares"
	"logistica/app/routes"
	"logistica/app/utility"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/mustache/v2"
)

var store = *session.New()
var client = *resty.New()

func main() {
	utility.ClearScreen()

	engine := mustache.New("./views", ".mustache")
	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Redirect("/404")
		},
	})

	app.Static("/", "./public")
	// Middleware global untuk menonaktifkan caching
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("noCache") == "true"
		},
		Expiration:   0 * time.Nanosecond,
		CacheControl: true,
	}))

	// Middlewares harus sebelum (Routes)
	app.Use(middlewares.UserAuthorization(&store))
	app.Use(middlewares.PathHandler(&store))

	// Routes
	routes.SetupRoutes(app, &store, &client)
	app.Get("/metrics", monitor.New())

	// Menjalankan server pada port 3000
	log.Fatal(app.Listen(":2500"))
}
