package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	rootArg := "."
	if len(os.Args) > 1 {
		rootArg = os.Args[1]
	}

	workspace, err := NewWorkspace(rootArg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Artemis v0.1.0 Coding Agent")
	fmt.Println("Workspace: %s\n", workspace.Root)
	fmt.Println("Type 'exit' or 'quit' to leave.")
	fmt.Println(strings.Repeat("-", 50))

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("artemis> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		if input == "exit" || input == "quit" {
			fmt.Println("bye bye!")
			break
		}

		fmt.Printf("prompt recieved %s\n", input)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}

}
