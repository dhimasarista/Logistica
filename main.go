package main

import (
	"log"
	"logistica/app/middlewares"
	"logistica/app/routes"
	"logistica/app/utility"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/mustache/v2"
)

var store = *session.New(session.ConfigDefault)
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
	app.Use(func(c *fiber.Ctx) error {
		if err := c.Next(); err != nil {
			c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			return err
		}
		return nil
	})
	app.Static("/", "./public")
	// Middleware global untuk menonaktifkan caching
	// app.Use(cache.New(cache.Config{
	// 	Next: func(c *fiber.Ctx) bool {
	// 		return c.Query("noCache") == "true"
	// 	},
	// 	Expiration:   0 * time.Nanosecond,
	// 	CacheControl: true,
	// }))

	// Middlewares harus sebelum (Routes)
	app.Use(middlewares.PathHandler(&store))
	app.Use(middlewares.UserAuthorization(&store))

	// Routes
	routes.SetupRoutes(app, &store, &client)
	app.Get("/metrics", monitor.New())

	// Menjalankan server pada port 3000
	log.Fatal(app.Listen(":6500"))
}
