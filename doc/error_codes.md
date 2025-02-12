# Error Codes

The events API will return various errors codes which can be used to identify 
the cause of the error. The error codes are defined in this file. Many of the codes
follow the HTTP status codes, but some are specific to the events API.

Furthermore, the use of the codes is specific to the events API and may not be 
a one-to-one match of the HTTP codes. Reading the description of the error code
is important to understand the cause of the error.

## Error Codes

#### 401 Unauthorized

This error indicates that the client is not authorized to perform the requested action.
This error can occur if the client has not authenticated with the server. This error does 
**not** occur if the client has authenticated but does not have the correct permissions, see 
the [403 Forbidden](#403-forbidden) error code.

##### Reasons

- **Not Authenticated**: The client has not successfully authenticated with the server.
- **Invalid Certificate**: The client has provided an invalid certificate.

<br>

#### 403 Forbidden

This error indicates that the client does not have the correct permissions to perform the
requested action. This error can only occur if the client has authenticated with the server 
but does not have the correct permissions. If the client has not authenticated, see the 
[401 Unauthorized](#401-unauthorized) error code.

##### Reasons

- **Insufficient Permissions**: The client does not have the correct permissions 
to perform the action.


#### 504 Service Unavailable

This error indicates that the services requested is not available. This error can occur if the
server is down or if the service is not available. This error can also occur if the server is
overloaded and cannot handle the request. This error can not be seen by a client if they are 
not authenticated with the server, see the [401 Unauthorized](#401-unauthorized) error code.

##### Reasons

- **Server Full**: The server has reached its maximum connection limit and cannot accept any 
more connections.

