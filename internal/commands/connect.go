package commands

import (
	"fmt"
	"opcli/internal/client"
)

// Connect подключается к OPC UA серверу по указанному endpoint
func Connect(endpoint string) error {
	if endpoint == "" {
		return fmt.Errorf("endpoint cannot be empty")
	}
	return client.Connect(endpoint)
}
