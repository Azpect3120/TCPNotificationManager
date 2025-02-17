package client

import (
	"fmt"
	"reflect"

	"github.com/Azpect3120/TCPNotificationManager/internal/events"
	"github.com/Azpect3120/TCPNotificationManager/internal/logger"
)

// HandleMessage is a function that will handle the message that is sent to the client.
// Very similar to how the server parses the events, but had to be abstracted here to
// keep the client code clean and easy to read.
//
// msg should be a byte slice that is sent from the server to the client. If the message
// is not a valid event, then an error will be thrown.
func (c *TcpClient) HandleMessage(msg []byte) {
	// Print the message to the client's logger, for debugging purposes.
	c.Logger.Log(string(msg)+"\n", logger.DEBUG)

	event, err := events.Parser(msg)
	if err != nil {
		// This happens when an event that is not implemented is received.
		c.Logger.Log(fmt.Sprintf("Error parsing message: %v\n", err), logger.ERROR)
		return
	}

	// This next section was copied from the server's HandleConnection method, it is
	// very hard to read and understand, but it makes the event creation and handling
	// pretty simple.
	//
	// All I have to do is register the event in the client's event handlers, and then
	// the client will handle the event when it is received.

	// Handle the event. A check for authorization should be done
	// in the handlers for the events, because there is no way to
	// get data from the raw interface{} type until it has been
	// type asserted.
	eventType := reflect.TypeOf(event).Elem()
	eventName := eventType.Name()

	if handler, ok := c.EventHandlers[eventName]; ok {
		// Correctly assert the handler type
		handlerType := reflect.TypeOf(handler)

		// Check if the handler type matches the event type. This
		// shit is black magic, Gemini created it for me, it seems
		// make sense but definitely not something I could have done
		// on my own.
		if handlerType.NumIn() == 2 && handlerType.In(1) == reflect.PointerTo(reflect.TypeOf(event).Elem()) {
			reflect.ValueOf(handler).Call([]reflect.Value{reflect.ValueOf(c), reflect.ValueOf(event)})
		} else {
			c.Logger.Log(fmt.Sprintln("Handler type mismatch for", eventName), logger.ERROR)
		}
	} else {
		c.Logger.Log(fmt.Sprintln("No handler found for", eventName), logger.ERROR)
	}
}
