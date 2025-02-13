package events

import "time"

// Base event structure which every event will inherit from.
// It is assumed that each event will have these details, so
// do not include specific event details here.
//
// Timestamp is using time.Time type, but an int64 might
// be more appropriate here.
type BaseEvent struct {
	Event     string    `json:"event"`
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

// Stores the content that should be inside the event.
type ConnectionAcceptedContent struct {
	ClientID string `json:"client_id"`
}

// Event returned by the server to the client when the connection
// is accepted.
type ConnectionAcceptedEvent struct {
	BaseEvent
	Content ConnectionAcceptedContent `json:"content"`
}

// Stores the content that should be inside the event.
type ConnectionRejectedContent struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
}

// Event returned by the server to the client when the connection
// is rejected.
type ConnectionRejectedEvent struct {
	BaseEvent
	Content ConnectionRejectedContent `json:"content"`
}

// Stores the content that should be inside the event.
//
// For now, this is to be ignored.
type RequestAuthenticationContent struct {
	Token string `json:"token"`
}

// Event sent by the client to the server when the connection
// is established.
type RequestAuthenticationEvent struct {
	BaseEvent
	Content RequestAuthenticationContent `json:"content"`
}

// Stores the content that should be inside the event.
type ClientAuthenticatedContent struct {
	ClientID string `json:"client_id"`
}

// Event sent by the server to the client when a new client
// authenticates with the server.
type ClientAuthenticatedEvent struct {
	BaseEvent
	Content ClientAuthenticatedContent `json:"content"`
}
