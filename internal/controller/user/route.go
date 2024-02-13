package controller

import (
	"github.com/gofiber/fiber/v2"
	"gridfs-file-manager/pkg/middlewares"
)

func (ctr *Controller) SetupRoute(app *fiber.App) {
	authRoute := app.Group("/api/user")
	authRoute.Get("/me", middlewares.Protected(), ctr.GetMe)
	authRoute.Get("/:userId", ctr.GetUser)

	//fileRoute := app.Group("/api/file")
	//fileRoute.Get("/shared", middlewares.Protected(), ctr.GetSharedFiles)
}
