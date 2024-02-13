package controller

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gridfs-file-manager/internal/services/file"
	"gridfs-file-manager/pkg/utils"
)

type FileService struct {
	services.FileManagement
}

type Controller struct {
	FileService
}

func NewController(fs FileService) *Controller {
	return &Controller{
		FileService: fs,
	}
}

func (ctr *Controller) DownloadOpenFile(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketId := c.Params("bucketId")

	file, filename, err := ctr.FileManagement.DownloadOpenFile(ctx, bucketId)
	if err != nil {
		return err
	}
	c.Set("Content-Disposition", "attachment; filename="+filename)
	return c.Send(file.Bytes())
}

func (ctr *Controller) DownloadPrivateFile(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketId := c.Params("bucketId")
	userId := c.Locals("userId").(string)

	file, filename, err := ctr.FileManagement.DownloadPrivateFile(ctx, bucketId, userId)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}
	c.Set("Content-Disposition", "attachment; filename="+filename)
	return c.Send(file.Bytes())
}

func (ctr *Controller) DownloadSharedFile(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketId := c.Params("bucketId")
	userId := c.Locals("userId").(string)

	file, filename, err := ctr.FileManagement.DownloadPrivateFile(ctx, bucketId, userId)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}
	c.Set("Content-Disposition", "attachment; filename="+filename)
	return c.Send(file.Bytes())
}

func (ctr *Controller) UploadFile(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	// parse request multipart form data and extract file
	file, err := c.FormFile("file")

	isPublic := c.Query("isPublic")

	if isPublic != "true" && isPublic != "false" {
		return utils.BadRequestErrorResponse(c, "Excepted isPublic true or false")
	}

	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	bucketId, err := ctr.FileManagement.UploadFile(file, userId, true)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	return utils.CreatedResponse(c, "Uploaded", fiber.Map{"bucketId": bucketId})
}

func (ctr *Controller) GetPublicFiles(c *fiber.Ctx) error {
	ctx := context.Background()
	files, err := ctr.FileManagement.GetPublicFiles(ctx)

	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	return utils.OkResponse(c, "Get", fiber.Map{"files": files})
}

func (ctr *Controller) GetPrivateFiles(c *fiber.Ctx) error {
	ctx := context.Background()
	userId := c.Locals("userId").(string)
	files, err := ctr.FileManagement.GetPrivateFiles(ctx, userId)

	if err != nil {
		return utils.BadRequestErrorResponse(c, fmt.Sprintf("ssas err %v", err))
	}

	return utils.OkResponse(c, "Get", fiber.Map{"files": files})
}

func (ctr *Controller) GetSharedFiles(c *fiber.Ctx) error {
	ctx := context.Background()
	userId := c.Locals("userId").(string)
	files, err := ctr.FileManagement.GetSharedFiles(ctx, userId)

	if err != nil {
		return utils.BadRequestErrorResponse(c, fmt.Sprintf("sas err %v", err))
	}

	return utils.OkResponse(c, "Get", fiber.Map{"files": files})
}

func (ctr *Controller) UpdateSharedFiles(c *fiber.Ctx) error {
	ctx := context.Background()

	var body struct {
		Files []string `json:"files" bson:"files"`
		Users []string `json:"users" bson:"users"`
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	err := ctr.FileManagement.UpdateSharedFiles(ctx, body.Users, body.Files)
	if err != nil {
		return err
	}

	return utils.OkResponse(c, "successfully update", fiber.Map{"files": body.Files, "usernames": body.Users})
}
