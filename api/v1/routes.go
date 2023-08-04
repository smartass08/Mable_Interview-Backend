package v1

import (
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(router fiber.Router, controller Controller) {
	AddFacebookRoutes(router, controller)
}

func AddFacebookRoutes(router fiber.Router, controller Controller) {
	router.Get(
		"/facebook/post",
		controller.SendEventData,
	)
}
