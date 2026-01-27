package client

import (
	"context"
	"fmt"

	"github.com/gopcua/opcua"
)

var client *opcua.Client

// Connect устанавливает соединение с OPC UA сервером
func Connect(endpoint string) error {
	if client != nil {
		fmt.Println("Already connected. Disconnecting first.")
		Disconnect()
	}

	fmt.Printf("Connecting to %s...\n", endpoint)

	ctx := context.Background()
	var err error
	client, err = opcua.NewClient(endpoint)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	if err := client.Connect(ctx); err != nil {
		client = nil
		return fmt.Errorf("failed to connect: %w", err)
	}

	fmt.Println("Successfully connected!")
	return nil
}

// Disconnect закрывает соединение с сервером
func Disconnect() error {
	if client != nil {
		fmt.Println("Disconnecting...")
		client.Close(context.Background())
		client = nil
	}
	return nil
}

// GetClient возвращает активный клиент или nil
func GetClient() *opcua.Client {
	return client
}
