# Events 

<!--toc:start-->
- [Events](#events)
  - [Base Event Structure](#base-event-structure)
  - [Server Sent Events](#server-sent-events)
    - [Accepted Connection](#accepted-connection)
    - [Refused Connection](#refused-connection)
    - [Client Authenticated](#client-authenticated)
    - [Client Disconnected](#client-disconnected)
  - [Client Sent Events](#client-sent-events)
    - [Request Authentication](#request-authentication)
    - [Disconnecting](#disconnecting)
<!--toc:end-->


## Base Event Structure

All events will have a standard structure. Events will also contain specific data based on the event type.

```json
{
    "event": "[event_name]",
    "id": "[sender_id]",
    "content": {
        "[key]": "[value]",
        ...
    },
    "timestamp": "[timestamp]"
}
```

## Server Sent Events

### Accepted Connection
When a client attempts to connect to the server, the server will either accept or refuse the connection. 
If the connection is accepted, the server will send an `connection_accepted` event to the client.

The `id` field will contain the ID of the server and the data will contain the client's new ID. The ID will
be generated by the server. The client does not need to do any work, just receive the message and save
the ID provided.

```json
{
    "event": "connection_accepted",
    "id": "[server_id]",
    "content": {
        "client_id": "[client_id]",
    },
    "timestamp": "[timestamp]"
}
```


### Refused Connection

When a client attempts to connect to the server, the server will either accept or refuse the connection. 
If the connection is accepted, the server will send an `connection_refused` event to the client.

The `id` field will contain the ID of the server and the data will contain the reason for the refusal, and
a status code to identify the error.

```json
{
    "event": "connection_refused",
    "id": "[server_id]",
    "content": {
        "code": "[code]",
        "reason": "[reason]",
    },
    "timestamp": "[timestamp]"
}
```

For details on the status codes and reasons, see the [Error Codes](error_codes.md) page.

### Client Authenticated

When a client authenticates with the server, the server will send a `client_authenticated` event to all other 
clients connected to the server. This event will contain the ID of the client that connected. Like all other 
messages, only authenticated clients will receive this message.

Originally, this was a client connected event, but it was changed to authenticated because the client is not
truly connected until they have authenticated. Plus, until the client is authenticated, the server knows nothing
about the client, other than the fact that they are trying to connect.

This message will not be sent back to the same client that authenticated, that would be silly.

```json
{
    "event": "client_authenticated",
    "id": "[server_id]",
    "content": {
        "client_id": "[client_id]",
    },
    "timestamp": "[timestamp]"
}
```

### Client Disconnected

When a client disconnects from the server, the server will send a `client_disconnected` event to all other clients
connected to the server. This event will contain the ID of the client that disconnected. The server will only send
this event when a client that was **authenticated** disconnects. If a client that was not authenticated disconnects,
the server will not send this event, however, it will still log the disconnection.

Similar to the `client_authenticated` event, this event will not be sent back to the client that disconnected. For
two reasons: 1) it would be stupid, 2) the client is disconnected, so they can't receive the message anyway.

```json
{
    "event": "client_disconnected",
    "id": "[server_id]",
    "content": {
        "client_id": "[client_id]",
    },
    "timestamp": "[timestamp]"
}
```

### Broadcast Message

When a client sends a message to the server, the server will broadcast that message to all other clients connected.
This event will contain the content and the sender of the message. The server will not send the message back to the 
sender, that would be silly.

Like all other events, only authenticated clients will receive this broadcast.

```json
{
    "event": "broadcast_message",
    "id": "[server_id]",
    "content": {
        "message": "[message]",
        "sender": "[client_id]"
    },
    "timestamp": "[timestamp]"
}
```


## Client Sent Events

### Request Authentication

This event is sent by a client when they first connect and need to authenticate with the server.
Nothing can be published or subscribed to until the client has been authenticated, so this event
is critical to the lifecycle of the client.

This request will contain basically nothing, just the event name. This is because in order to connect
to the TCP server to begin with, the keys and certificates must be provided. The server will use these
to assume the client is authenticated. 

The ID field will contain nothing for this event because the server generates the ID for the client,
however due to the inherent nature of the event, the ID field is required. The content field will 
contain a token that the client must provide to authenticate.

**TODO:** Implement tokens for this event to add another layer of security. For now, token can be ignored.

```json
{
    "event": "request_authentication",
    "id": "__",
    "content": {
        "token": "[token]"
    },
    "timestamp": "[timestamp]"
}
```

### Disconnecting

When a client wants to disconnect from the server, they will send a `disconnecting` event to the server.
This event will contain the ID of the client that is disconnecting. The server will then remove the client
from the list of connected clients and send a `client_disconnected` event to all other clients.

```json
{
    "event": "disconnecting",
    "id": "[client_id]",
    "content": {},
    "timestamp": "[timestamp]"
}
```

### Send Message

When a client wants to send a message to the server, they will send a `send_message` event to the server.
This message can be anything, and the server will simply broadcast the message to all other clients connected
and authenticated. This can be used as a simple messaging system, once a UI has been implemented.

```json
{
    "event": "send_message",
    "id": "[client_id]",
    "content": {
        "message": "[message]"
    },
    "timestamp": "[timestamp]"
}
```
