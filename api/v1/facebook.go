package v1

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"demo/logging"
	"demo/schemas"
	"demo/service"
	serviceschemas "demo/service/schemas"
)

func (Controller) SendEventData(ctx *fiber.Ctx) error {
	logger := logging.GetLogger()

	// Parse the request body into a struct
	var request schemas.FacebookEventPayload
	err := ctx.BodyParser(&request)
	if err != nil {
		logger.Error("Error occurred while parsing request body", zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.ErrorResponse{Detail: err.Error()})
	}

	err = request.Validate()
	if err != nil {
		logger.Error("Error occurred while validating request body", zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.ErrorResponse{Detail: err.Error()})
	}

	// Create a request body for events api
	var requestPayload = make([]serviceschemas.Event, 0)
	for _, event := range request.Events {
		var eventPayload serviceschemas.Event
		eventPayload.User.Initialise()

		// Create a map of custom data
		customData := map[string]string{
			"currency": event.SessionData.IPLocationData.Currency,
			"value":    fmt.Sprintf("%.2f", event.SessionData.CartData.TotalPrice),

			// Add more key-value pairs as needed based on the other fields in the event struct
			// For example: "order_id": fmt.Sprintf("%d", event.SourceCustomData.Shopify.OrderID),
			// Add other fields as needed.
		}

		if event.SessionData.CustomerData.Email != "" {
			// Hash the email address
			eventPayload.User.Email = append(eventPayload.User.Email, hashString(event.SessionData.CustomerData.Email))
		}

		if event.SessionData.CustomerData.Phone != "" {
			// Hash the phone number
			eventPayload.User.Phone = append(eventPayload.User.Phone, hashString(event.SessionData.CustomerData.Phone))
		}

		eventPayload.EventName = event.EventName
		eventPayload.EventTime = time.Now().Unix()
		eventPayload.CustomData = customData

		requestPayload = append(requestPayload, eventPayload)
	}

	// Send the request payload to the service
	// Create facebook service client
	facebookServiceClient := service.FaceBookConversationAPIClient{
		BaseURL:   "https://graph.facebook.com/v11.0/",
		PixelID:   request.DestinationSettings.PixelID,
		AccessKey: request.DestinationSettings.AccessKey,
	}

	err = facebookServiceClient.SendEvents(request.DestinationSettings.TestEventCode, requestPayload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schemas.ErrorResponse{Detail: err.Error()})
	}

	// Create an output message
	var output = schemas.FacebookEventOutput{
		Message: "Successfully sent the following data to the service",
		Events:  requestPayload,
	}

	return ctx.JSON(output)
}
