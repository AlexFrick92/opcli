package client

import (
	"context"
	"fmt"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
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

	// Получаем и выводим информацию о сервере
	info, err := GetServerInfo()
	if err != nil {
		fmt.Printf("Warning: could not retrieve server info: %v\n", err)
	} else {
		printServerInfo(info)
	}

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

// ServerInfo содержит информацию о OPC UA сервере
type ServerInfo struct {
	ProductName      string
	ManufacturerName string
	SoftwareVersion  string
	ServerState      string
}

// GetServerInfo получает информацию о сервере
func GetServerInfo() (*ServerInfo, error) {
	if client == nil {
		return nil, fmt.Errorf("not connected to server")
	}

	ctx := context.Background()

	// Node IDs из стандартного адресного пространства OPC UA
	productNameID := "i=2261"
	manufacturerID := "i=2262"
	versionID := "i=2263"
	stateID := "i=2259"

	info := &ServerInfo{}

	// Читаем ProductName
	if val, err := readNodeValue(ctx, productNameID); err == nil {
		info.ProductName = val
	}

	// Читаем ManufacturerName
	if val, err := readNodeValue(ctx, manufacturerID); err == nil {
		info.ManufacturerName = val
	}

	// Читаем SoftwareVersion
	if val, err := readNodeValue(ctx, versionID); err == nil {
		info.SoftwareVersion = val
	}

	// Читаем ServerState
	if val, err := readNodeValue(ctx, stateID); err == nil {
		info.ServerState = val
	}

	return info, nil
}

// printServerInfo выводит информацию о сервере в консоль
func printServerInfo(info *ServerInfo) {
	fmt.Println("\n=== Server Information ===")
	if info.ProductName != "" {
		fmt.Printf("Product:      %s\n", info.ProductName)
	}
	if info.ManufacturerName != "" {
		fmt.Printf("Manufacturer: %s\n", info.ManufacturerName)
	}
	if info.SoftwareVersion != "" {
		fmt.Printf("Version:      %s\n", info.SoftwareVersion)
	}
	if info.ServerState != "" {
		fmt.Printf("State:        %s\n", info.ServerState)
	}
	fmt.Println("==========================")
}

// readNodeValue читает значение узла по Node ID
func readNodeValue(ctx context.Context, nodeID string) (string, error) {
	id, err := ua.ParseNodeID(nodeID)
	if err != nil {
		return "", fmt.Errorf("invalid node ID: %w", err)
	}

	req := &ua.ReadRequest{
		MaxAge:             2000,
		NodesToRead:        []*ua.ReadValueID{{NodeID: id}},
		TimestampsToReturn: ua.TimestampsToReturnBoth,
	}

	resp, err := client.Read(ctx, req)
	if err != nil {
		return "", fmt.Errorf("read failed: %w", err)
	}

	if len(resp.Results) == 0 {
		return "", fmt.Errorf("no results")
	}

	result := resp.Results[0]
	if result.Status != ua.StatusOK {
		return "", fmt.Errorf("bad status: %v", result.Status)
	}

	return fmt.Sprintf("%v", result.Value.Value()), nil
}
