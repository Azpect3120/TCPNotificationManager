package server

import (
	"fmt"
	"net"

	"github.com/Azpect3120/TCPNotificationManager/internal/events"
)

// Important code segment
// Switch on the event type and handle the event
// This is temporary until the event handlers are implemented.
// switch e := event.(type) {
// case *RequestAuthenticationEvent:
// 	fmt.Printf("RequestAuthenticationEvent: %s\n", e.BaseEvent.Event)
// case *ConnectionAcceptedEvent:
// 	fmt.Printf("ConnectionAcceptedEvent: %s\n", e.BaseEvent.Event)
// case *ConnectionRejectedEvent:
// 	fmt.Printf("ConnectionRejectedEvent: %s\n", e.BaseEvent.Event)
// default:
// 	fmt.Printf("Unknown event type: %v\n", e)
// }

func RequestAuthenticationHandler(conn net.Conn, event *events.RequestAuthenticationEvent) {
	fmt.Printf("%+v\n", event)
}
