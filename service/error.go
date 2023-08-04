package service

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"demo/logging"
)

// HandleError handles unhandled errors
func HandleError(ctx *fiber.Ctx, err error) error {
	logger := logging.GetLogger()
	code := http.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	} else {
		logger.Error("Error occurred", zap.Error(err))
	}

	return ctx.Status(code).JSON(fiber.Map{"detail": err.Error()})
}
