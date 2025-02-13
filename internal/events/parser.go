package events

import (
	"encoding/json"
	"fmt"
)

// Parse the event type from the data. This is a complex problem due to the
// polymorphic nature of the events. The event type is stored in the "event"
// field of the JSON data, and can therefore be marshalled into a struct.
//
// The event type will be returned as an interface{} and will need to be
// asserted to the correct type.
func _eventType(data []byte) (interface{}, error) {
	// Use a struct to unmarshal the event type.
	var eventType struct {
		Event string `json:"event"`
	}
	if err := json.Unmarshal(data, &eventType); err != nil {
		return nil, fmt.Errorf("failed to determine event type: %w", err)
	}

	var event interface{}
	switch eventType.Event {
	case "connection_accepted":
		event = &ConnectionAcceptedEvent{}
	case "connection_rejected":
		event = &ConnectionRejectedEvent{}
	case "request_authentication":
		event = &RequestAuthenticationEvent{}
	default:
		return nil, fmt.Errorf("Event type '%s' has not been implemented.", eventType.Event)
	}

	return event, nil
}

// Parser is used to parse the data received from the server/client. This
// function will determine the event type and unmarshal the data into the
// event. The event will then be handled by the caller. The map stored in
// in the client or server should be used to map a function to the type of
// event.
//
// The event type will be returned as an interface{} and will be used
// in the server package to run the event handlers.
//
// Events are not scoped to server or client so this function can be used
// for both the server and the client. However, do not check for server
// events on the client, or vice versa.
func Parser(data []byte) (interface{}, error) {
	event, err := _eventType(data)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, event); err != nil {
		return nil, err
	}

	return event, nil
}
