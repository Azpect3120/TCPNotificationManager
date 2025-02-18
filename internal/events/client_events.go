package events

import "time"

// Create and return a new RequestAuthenticationEvent. This function does not
// generate any details, instead it requires all details as arguments. Which
// should be generated elsewhere.
//
// For now, token can be ignored and anything can be passed. This is temporary
// until token authentication is implemented.
//
// All timestamps will be sent back in UTC format.
func NewRequestAuthenticationEvent(token string) RequestAuthenticationEvent {
	return RequestAuthenticationEvent{
		BaseEvent: BaseEvent{
			Event:     "request_authentication",
			ID:        "",
			Timestamp: time.Now().UTC(),
		},
		Content: RequestAuthenticationContent{
			Token: token,
		},
	}
}

// Create and return a new ClientDisconnectingEvent. This function does not
// generate any details, instead it requires all details as arguments. Which
// should be generated elsewhere.
//
// For now, token can be ignored and anything can be passed. This is temporary
// until token authentication is implemented.
//
// The content for this event is blank so a default EmptyContent struct is
// used to fill the content field.
//
// All timestamps will be sent back in UTC format.
func NewClientDisconnectingEvent(clientID string) ClientDisconnectingEvent {
	return ClientDisconnectingEvent{
		BaseEvent: BaseEvent{
			Event:     "disconnecting",
			ID:        clientID,
			Timestamp: time.Now().UTC(),
		},
		Content: EmptyContent{},
	}
}

// Create and return a new ClientDisconnectingEvent. This function does not
// generate any details, instead it requires all details as arguments. Which
// should be generated elsewhere.
//
// The message should be a complete string, nothing will be done in this function
// to ensure that the message is valid, or formatted.
//
// All timestamps will be sent back in UTC format.
func NewSendMessageEvent(clientID, message string) SendMessageEvent {
	return SendMessageEvent{
		BaseEvent: BaseEvent{
			Event:     "send_message",
			ID:        clientID,
			Timestamp: time.Now().UTC(),
		},
		Content: SendMessageContent{
			Message: message,
		},
	}
}
