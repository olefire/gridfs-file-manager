package controller

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gridfs-file-manager/internal/models"
	"gridfs-file-manager/internal/services/auth"
	"gridfs-file-manager/pkg/utils"
)

type AuthService struct {
	services.AuthManagement
}

type Controller struct {
	AuthService
}

func NewController(as AuthService) *Controller {
	return &Controller{
		AuthService: as,
	}
}

func (ctr *Controller) Signup(c *fiber.Ctx) error {
	ctx := context.Background()
	var signupBody models.SignupBody

	if err := c.BodyParser(&signupBody); err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}
	// validate user input
	errors := utils.ValidateStruct(signupBody)
	if errors != nil {
		return utils.UnprocessedInputResponse(c, fiber.Map{
			"errors": errors,
		})
	}

	insertedUser, err := ctr.AuthManagement.SignUpUser(ctx, &signupBody)
	if err != nil {
		return utils.BadRequestErrorResponse(c, fmt.Sprint(err))
	}

	return utils.CreatedResponse(c, "Account created", fiber.Map{
		"userId": insertedUser,
	})
}

func (ctr *Controller) Login(c *fiber.Ctx) error {
	ctx := context.Background()
	var loginBody models.LoginRequest

	if err := c.BodyParser(&loginBody); err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	// validate users input
	errors := utils.ValidateStruct(loginBody)
	if errors != nil {
		return utils.UnprocessedInputResponse(c, fiber.Map{"errors": errors})
	}

	accessToken, err := ctr.AuthService.SignInUser(ctx, &loginBody)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	return utils.OkResponse(c, "Login successfully", fiber.Map{
		"accessToken": accessToken,
	})
}
