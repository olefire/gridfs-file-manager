package controller

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gridfs-file-manager/internal/services/user"
	"gridfs-file-manager/pkg/utils"
)

type UserService struct {
	services.UserManagement
}

type Controller struct {
	UserService
}

func NewController(us UserService) *Controller {
	return &Controller{
		UserService: us,
	}
}

func (ctr *Controller) GetUser(c *fiber.Ctx) error {
	ctx := context.Background()
	userId := c.Params("userId")

	user, err := ctr.UserManagement.GetUser(ctx, userId)
	if err != nil {
		return utils.NotFoundErrorResponse(c)
	}

	return utils.OkResponse(c, "Get user details", fiber.Map{
		"user": user,
	})
}

func (ctr *Controller) GetMe(c *fiber.Ctx) error {
	ctx := context.Background()
	userId := c.Locals("userId").(string)

	user, err := ctr.UserManagement.GetUser(ctx, userId)
	if err != nil {
		return utils.BadRequestErrorResponse(c, fmt.Sprintf("can`t find user: %v", err))
	}

	return utils.OkResponse(c, "Get current user", fiber.Map{
		"user": user,
	})
}

//func (ctr *Controller) GetSharedFiles(c *fiber.Ctx) error {
//	ctx := context.Background()
//	userId := c.Locals("userId").(string)
//
//	files, err := ctr.UserManagement.GetSharedFiles(ctx, userId)
//	if err != nil {
//		return utils.InternalServerErrorResponse(c, err)
//	}
//
//	return utils.OkResponse(c, "Get shared files", fiber.Map{
//		"shared files": files,
//	})
//}
