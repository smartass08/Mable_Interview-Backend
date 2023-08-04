package schemas

import serviceschemas "demo/service/schemas"

type FacebookEventOutput struct {
	Message string                 `json:"message,omitempty"`
	Events  []serviceschemas.Event `json:"events,omitempty"`
}
