package pkg

import (
	"github.com/gofiber/fiber/v2"
	error2 "simple-start-page/internal/error"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func GenerateSuccessResponse(ctx *fiber.Ctx, data interface{}) error {
	return ctx.JSON(Response{
		Code:    0,
		Message: "OK",
		Data:    data,
	})
}

func GenerateErrorResponse(ctx *fiber.Ctx, err error) error {
	if appErr, ok := err.(*error2.AppError); ok {
		switch appErr.Code {
		case error2.ErrCodeInvalidCredential:
			ctx.Status(400)
		case error2.ErrCodeUnauthenticated:
			ctx.Status(401)
		case error2.ErrCodeValidationFailed:
			ctx.Status(400)
		default:
			ctx.Status(500)
		}
		return ctx.JSON(appErr)
	}
	return wrapGeneralError(ctx, err)
}

func wrapGeneralError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(500).JSON(&Response{
		Code:    int(error2.ErrCodeInternal),
		Message: err.Error(),
	})
}
