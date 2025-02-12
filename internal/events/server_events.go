package events

import "time"

// Create and return a new ConnectionAcceptedEvent. This function does not
// generate any details, instead it requires all details as arguments. Which
// should be generated elsewhere.
func NewConnectionAcceptedEvent(serverID, clientID string) ConnectionAcceptedEvent {
	return ConnectionAcceptedEvent{
		BaseEvent: BaseEvent{
			Event:     "connection_accepted",
			ID:        serverID,
			Timestamp: time.Now(),
		},
		Content: ConnectionAcceptedContent{ClientID: clientID},
	}
}
