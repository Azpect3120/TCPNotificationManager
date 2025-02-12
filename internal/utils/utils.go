package utils

import (
	"fmt"

	"github.com/google/uuid"
)

// Create a random ID for the server. This is abstracted for
// future use in generating more specific/complex server IDs.
func GenerateServerID() string {
	return fmt.Sprintf("server-%s", uuid.NewString())
}

// Create a random ID for the client. This is abstracted for
// future use in generating more specific/complex client IDs.
func GenerateClientID() string {
	return fmt.Sprintf("client-%s", uuid.NewString())
}
