package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     []byte("SECRET"),
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
		ContextKey:     "jwtToken",
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "missing or malformed JWT" {
		err := c.Next()
		if err != nil {
			return err
		}
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func jwtSuccess(ctx *fiber.Ctx) error {
	token := ctx.Locals("jwtToken").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	strId := claims["userId"].(string)

	ctx.Locals("userId", strId)

	return ctx.Next()
}
