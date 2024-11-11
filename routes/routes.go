package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/codingsluv/book-store/POS-api/controller"
)

func SetupRoute(app *fiber.App) {
	// Authentication
	app.Get("/", controller.Login)

	// Cashier
	app.Get("/cashier", controller.CashierList)
	app.Post("/cashier", controller.CreateCashier)
	app.Get("/cashier/:cashierID", controller.GetCashier)
	app.Put("/cashier/:cashierID", controller.UpdateCashier)
	app.Delete("/cashier/:cashierID", controller.DeleteCashier)
}
