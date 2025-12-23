package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("opcli - OPC UA Interactive Client")
	fmt.Println("Type 'help' for available commands")
	fmt.Println()

	// Если передан аргумент connect, выполняем сразу
	if len(os.Args) > 2 && os.Args[1] == "connect" {
		endpoint := os.Args[2]
		fmt.Printf("Connecting to %s...\n", endpoint)
		// TODO: реализовать подключение
	}

	// Запуск интерактивной оболочки
	runShell()
}

func runShell() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("opcli> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if err := executeCommand(input); err != nil {
			if err.Error() == "exit" {
				fmt.Println("Goodbye!")
				return
			}
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	}
}

func executeCommand(input string) error {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil
	}

	command := parts[0]
	args := parts[1:]

	switch command {
	case "help":
		printHelp()
	case "connect":
		return handleConnect(args)
	case "browse":
		return handleBrowse(args)
	case "read":
		return handleRead(args)
	case "write":
		return handleWrite(args)
	case "subscribe":
		return handleSubscribe(args)
	case "call":
		return handleCall(args)
	case "exit", "quit":
		return fmt.Errorf("exit")
	default:
		return fmt.Errorf("unknown command: %s. Type 'help' for available commands", command)
	}

	return nil
}

func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  connect <endpoint>     - Connect to OPC UA server")
	fmt.Println("  browse [nodeId]        - Browse node tree structure")
	fmt.Println("  read <nodeId>          - Read node value")
	fmt.Println("  write <nodeId> <value> - Write value to node")
	fmt.Println("  subscribe <nodeId>     - Subscribe to node value changes")
	fmt.Println("  call <nodeId> [args]   - Call server method")
	fmt.Println("  help                   - Show this help")
	fmt.Println("  exit, quit             - Exit the program")
	fmt.Println()
	fmt.Println("NodeId format: ns=<namespace>;i=<identifier>")
	fmt.Println("Example: ns=2;i=1000")
}

func handleConnect(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: connect <endpoint>")
	}
	endpoint := args[0]
	fmt.Printf("Connecting to %s...\n", endpoint)
	// TODO: реализовать подключение через OPC UA client
	fmt.Println("TODO: Connection not implemented yet")
	return nil
}

func handleBrowse(args []string) error {
	fmt.Println("Browsing nodes...")
	// TODO: реализовать browse
	fmt.Println("TODO: Browse not implemented yet")
	return nil
}

func handleRead(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: read <nodeId>")
	}
	nodeID := args[0]
	fmt.Printf("Reading node %s...\n", nodeID)
	// TODO: реализовать read
	fmt.Println("TODO: Read not implemented yet")
	return nil
}

func handleWrite(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: write <nodeId> <value>")
	}
	nodeID := args[0]
	value := args[1]
	fmt.Printf("Writing value %s to node %s...\n", value, nodeID)
	// TODO: реализовать write
	fmt.Println("TODO: Write not implemented yet")
	return nil
}

func handleSubscribe(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: subscribe <nodeId>")
	}
	nodeID := args[0]
	fmt.Printf("Subscribing to node %s...\n", nodeID)
	// TODO: реализовать subscribe
	fmt.Println("TODO: Subscribe not implemented yet")
	return nil
}

func handleCall(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: call <nodeId> [args...]")
	}
	nodeID := args[0]
	methodArgs := args[1:]
	fmt.Printf("Calling method %s with args: %v\n", nodeID, methodArgs)
	// TODO: реализовать call
	fmt.Println("TODO: Method call not implemented yet")
	return nil
}
