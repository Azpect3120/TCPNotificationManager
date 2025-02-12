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
