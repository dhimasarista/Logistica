package main

import (
	"log"
	"logistica/app/middlewares"
	"logistica/app/routes"
	"os"
	"os/exec"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/template/mustache/v2"
)

func main() {
	clearScreen()

	engine := mustache.New("./views", ".mustache")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")

	// Middleware global untuk menonaktifkan caching
	app.Use(func(c *fiber.Ctx) error {
		// Mengatur header Cache-Control untuk menonaktifkan caching
		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		return c.Next()
	})

	// Rute untuk menampilkan halaman HTML
	routes.SetupRoutes(app)
	middlewares.UserAuthorization(app) // Menangani autorisasi user
	app.Get("/metrics", monitor.New())

	// Menjalankan server pada port 3000
	log.Fatal(app.Listen(":4500"))
}

func clearScreen() {
	// Menentukan perintah untuk membersihkan layar sesuai dengan sistem operasi yang digunakan
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
