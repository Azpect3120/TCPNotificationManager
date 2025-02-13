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

// Check if an item exists in a slice. This function is generic
// and can be used with any type that is comparable.
func Contains[T comparable](slice []T, item T) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
