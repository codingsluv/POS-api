package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/codingsluv/book-store/POS-api/config"
	"github.com/codingsluv/book-store/POS-api/routes"
)

func main() {
	fmt.Println("pos api code is running")
	// config.ConnectDB()
	config.ConnectDB()

	app := fiber.New()

	routes.SetupRoute(app)
	app.Listen(":3000")
}
