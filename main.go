package main

import (
	"log"
	"logistica/app/middlewares"
	"logistica/app/routes"
	"os"
	"os/exec"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
)

func main() {
	clearScreen()

	engine := mustache.New("./views", ".mustache")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")

	// Rute untuk menampilkan halaman HTML
	routes.IndexRoutes(app)
	routes.AuthRoutes(app)
	routes.DashboardRoutes(app)
	routes.DaftarSuratRoutes(app)
	routes.UnggahSuratRoutes(app)
	middlewares.UserAuthorization(app)

	// Menjalankan server pada port 3000
	log.Fatal(app.Listen(":3000"))
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
