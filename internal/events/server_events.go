package events

import "time"

// Create and return a new ConnectionAcceptedEvent. This function does not
// generate any details, instead it requires all details as arguments. Which
// should be generated elsewhere.
//
// All timestamps will be sent back in UTC format.
func NewConnectionAcceptedEvent(serverID, clientID string) ConnectionAcceptedEvent {
	return ConnectionAcceptedEvent{
		BaseEvent: BaseEvent{
			Event:     "connection_accepted",
			ID:        serverID,
			Timestamp: time.Now().UTC(),
		},
		Content: ConnectionAcceptedContent{
			ClientID: clientID,
		},
	}
}

// Create and return a new ConnectionRejectedEvent. This function does not
// generate any details, instead it requires all details as arguments. Which
// should be generated elsewhere.
//
// All timestamps will be sent back in UTC format.
func NewConnectionRejectedEvent(serverID string, code int, reason string) ConnectionRejectedEvent {
	return ConnectionRejectedEvent{
		BaseEvent: BaseEvent{
			Event:     "connection_rejected",
			ID:        serverID,
			Timestamp: time.Now().UTC(),
		},
		Content: ConnectionRejectedContent{
			Code:   code,
			Reason: reason,
		},
	}
}

// Create and return a new ClientAuthenticatedEvent. This function does not
// generate any details, instead it requires all details as arguments. Which
// should be generated elsewhere.
//
// All timestamps will be sent back in UTC format.
func NewClientAuthenticatedEvent(serverID, clientID string) ClientAuthenticatedEvent {
	return ClientAuthenticatedEvent{
		BaseEvent: BaseEvent{
			Event:     "client_authenticated",
			ID:        serverID,
			Timestamp: time.Now().UTC(),
		},
		Content: ClientAuthenticatedContent{
			ClientID: clientID,
		},
	}
}

// Create and return a new ClientDisconnectedEvent. This function does not
// generate any details, instead it requires all details as arguments. Which
// should be generated elsewhere.
//
// All timestamps will be sent back in UTC format.
func NewClientDisconnectedEvent(serverID, clientID string) ClientDisconnectedEvent {
	return ClientDisconnectedEvent{
		BaseEvent: BaseEvent{
			Event:     "client_disconnected",
			ID:        serverID,
			Timestamp: time.Now().UTC(),
		},
		Content: ClientDisconnectedContent{
			ClientID: clientID,
		},
	}
}

// Create and return a new BroadcastMessageEvent. This function does not
// generate any details, instead it requires all details as arguments. Which
// should be generated elsewhere.
//
// The message should be a complete string, nothing will be done in this function
// to ensure that the message is valid, or formatted.
//
// All timestamps will be sent back in UTC format.
func NewBroadcastMessageEvent(serverID, clientID, message string) BroadcastMessageEvent {
	return BroadcastMessageEvent{
		BaseEvent: BaseEvent{
			Event:     "broadcast_message",
			ID:        serverID,
			Timestamp: time.Now().UTC(),
		},
		Content: BroadcastMessageContent{
			Sender:  clientID,
			Message: message,
		},
	}
}
