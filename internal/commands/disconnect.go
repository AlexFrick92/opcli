package commands

import (
	"github.com/alexfrick92/opcli/internal/client"
)

// Disconnect отключается от OPC UA сервера
func Disconnect() error {
	return client.Disconnect()
}
