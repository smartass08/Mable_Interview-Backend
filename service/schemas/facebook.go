package schemas

type User struct {
	Email []string `json:"em,omitempty"`
	Phone []string `json:"ph,omitempty"`
}

func (u *User) Initialise() {
	u.Email = make([]string, 0)
	u.Phone = make([]string, 0)
}

type Event struct {
	EventName  string            `json:"event_name"`
	EventTime  int64             `json:"event_time"`
	User       User              `json:"user_data"`
	CustomData map[string]string `json:"custom_data"`
}

type FacebookEventPayload struct {
	Data          []Event `json:"data"`
	TestEventCode string  `json:"test_event_code,omitempty"`
}
