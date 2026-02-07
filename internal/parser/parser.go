package parser

import (
	"fmt"
	"net"
	"strings"

	"github.com/alexfrick92/opcli/internal/commands"
)

var connectCommand = commands.Connect
var disconnectCommand = commands.Disconnect

// Execute выполняет команду из пользовательского ввода
func Execute(input string) error {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil
	}

	command := parts[0]
	args := parts[1:]

	switch command {
	case "help":
		PrintHelp()
	case "connect":
		return handleConnect(args)
	case "disconnect":
		return handleDisconnect()
	case "exit", "quit":
		return fmt.Errorf("exit")
	default:
		return fmt.Errorf("unknown command: %s. Type 'help' for available commands", command)
	}

	return nil
}

// PrintHelp выводит справку по доступным командам
func PrintHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  connect <endpoint>  - Connect to OPC UA server")
	fmt.Println("  disconnect          - Disconnect from server")
	fmt.Println("  help                - Show this help")
	fmt.Println("  exit, quit          - Exit the program")
}

func handleConnect(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: connect <endpoint>")
	}
	return connectCommand(args[0])
}

func handleDisconnect() error {
	return disconnectCommand()
}

// ParseStartupArgs обрабатывает аргументы командной строки при запуске
func ParseStartupArgs(args []string) error {
	// Если передан IP-адрес, подключаемся с портом по умолчанию
	if len(args) == 2 && isIPv4(args[1]) {
		endpoint := fmt.Sprintf("opc.tcp://%s:4840", args[1])
		return connectCommand(endpoint)
	}

	// Handle 'connect' command
	if len(args) >= 2 && args[1] == "connect" {
		if len(args) < 3 { // 'connect' command requires an endpoint
			return fmt.Errorf("usage: connect <endpoint>")
		}
		return connectCommand(args[2])
	}

	return nil
}

func isIPv4(s string) bool {
	ip := net.ParseIP(s)
	return ip != nil && ip.To4() != nil
}
