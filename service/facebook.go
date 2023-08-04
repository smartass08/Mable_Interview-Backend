package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"go.uber.org/zap"

	"demo/logging"
	serviceschemas "demo/service/schemas"
)

type FaceBookConversationAPIClient struct {
	BaseURL   string
	PixelID   string
	AccessKey string
}

func (api *FaceBookConversationAPIClient) postEventsToFacebook(requestPayload []byte) (err error) {
	logger := logging.GetLogger()
	endpointURL, err := url.Parse(api.BaseURL + api.PixelID + "/events")
	if err != nil {
		logger.Error("Error parsing URL", zap.Error(err), zap.String("URL", endpointURL.String()))
		return
	}

	// Create the URL parameters
	params := url.Values{}
	params.Add("access_token", api.AccessKey)

	// Add the parameters to the URL
	endpointURL.RawQuery = params.Encode()

	// Make the POST request with the request payload
	resp, err := http.Post(endpointURL.String(), "application/json", bytes.NewBuffer(requestPayload))
	if err != nil {
		logger.Error("Error making POST request", zap.Error(err), zap.String("URL", endpointURL.String()))
		return
	}

	// Close the response body after the function returns
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.Error("Error closing response body", zap.Error(err))
		}
	}(resp.Body)

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		logger.Error("Error making POST request", zap.Error(err), zap.String("URL", endpointURL.String()))
		return
	}

	// Read and print the response body (optional)
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body", zap.Error(err))
		return
	}

	logger.Info("Facebook Events Response body", zap.String("body", string(responseBody)))
	return
}

func (api *FaceBookConversationAPIClient) SendEvents(testEventCode string, events []serviceschemas.Event) (err error) {
	logger := logging.GetLogger()
	logger.Info("Sending events to Facebook", zap.Any("events", events))
	// Create the request payload
	payload := serviceschemas.FacebookEventPayload{Data: events}

	// Check if test event code is present
	if testEventCode != "" {
		payload.TestEventCode = testEventCode
	}

	// Marshal the payload
	requestPayload, err := json.Marshal(payload)
	if err != nil {
		logger.Error("Error marshalling payload", zap.Error(err))
		return
	}

	// Make the POST request
	err = api.postEventsToFacebook(requestPayload)
	return
}
