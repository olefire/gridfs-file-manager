package controller

import (
	"github.com/gofiber/fiber/v2"
)

func (ctr *Controller) SetupRoute(app *fiber.App) {
	authRoute := app.Group("/api/auth")
	authRoute.Post("/signup", ctr.Signup)
	authRoute.Post("/login", ctr.Login)
}
