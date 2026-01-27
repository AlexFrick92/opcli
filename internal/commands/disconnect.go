package commands

import (
	"opcli/internal/client"
)

// Disconnect отключается от OPC UA сервера
func Disconnect() error {
	return client.Disconnect()
}
