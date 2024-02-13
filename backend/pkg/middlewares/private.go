package middlewares

//import (
//	"fmt"
//	"github.com/gofiber/fiber/v2"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"net/http"
//)
//
//func IsAvailableFile() fiber.Handler {
//	return func(ctx *fiber.Ctx) error {
//		currentUserId := ctx.Locals("userId")
//		curId := fmt.Sprint(currentUserId)
//
//		bucketId, err := primitive.ObjectIDFromHex(ctx.Params("id"))
//
//		if err != nil {
//			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": err.Error()})
//		}
//
//		paste, err := pasteService.GetPasteById(ctx, id)
//
//		if err != nil {
//			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
//			return
//		}
//
//		if paste.UserID.Hex() != curId {
//			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "you don't have permission to access this resource"})
//			err = fmt.Errorf("you don't have permission to access this resource")
//		}
//
//		if err != nil {
//			ctx.Abort()
//		}
//
//		ctx.Next()
//	}
//}
