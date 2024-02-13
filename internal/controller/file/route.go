package controller

import (
	"github.com/gofiber/fiber/v2"
	"gridfs-file-manager/pkg/middlewares"
)

func (ctr *Controller) SetupRoute(app *fiber.App) {
	fileRoute := app.Group("/api/media")
	fileRoute.Get("/:bucketId", ctr.DownloadOpenFile)
	fileRoute.Get("/protected/:bucketId", middlewares.Protected(), ctr.DownloadPrivateFile)
	fileRoute.Get("/shared/:bucketId", middlewares.Protected(), ctr.DownloadPrivateFile)
	fileRoute.Post("/", middlewares.Protected(), ctr.UploadFile)

	mediaRoute := app.Group("/api/file")
	mediaRoute.Get("/public", ctr.GetPublicFiles)
	mediaRoute.Get("/private", middlewares.Protected(), ctr.GetPrivateFiles)
	mediaRoute.Get("/shared", middlewares.Protected(), ctr.GetSharedFiles)
	mediaRoute.Patch("/share", middlewares.Protected(), ctr.UpdateSharedFiles)
}
