package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()
	setUpRoutes(app)
	app.Listen(3000)
}

func setUpRoutes(app *fiber.App) {

}
