package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alexfrick92/opcli/internal/client"
	"github.com/alexfrick92/opcli/internal/parser"
)

func main() {
	fmt.Println("opcli - OPC UA Interactive Client")
	fmt.Println("Type 'help' for available commands")
	fmt.Println()

	if err := parser.ParseStartupArgs(os.Args); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	runShell()
}

func runShell() {
	defer client.Disconnect()
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

		if err := parser.Execute(input); err != nil {
			if err.Error() == "exit" {
				fmt.Println("Goodbye!")
				return
			}
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	}
}
