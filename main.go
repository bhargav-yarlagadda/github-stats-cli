package main

import (
	"bufio"
	"fmt"
	"github-stats-cli/commands"
	"os"
	"strings"
)

func main() {
	// Initialize commands
	commands.InitCommands()
	fmt.Fprint(os.Stdout, "Welcome to GitHub CLI Tool \n\n\n")

	// Start command loop
	for {
		// Print the prompt
		fmt.Fprintf(os.Stdout, "$ghcli ")

		// Read the input line using a scanner
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // Read the next line

		// Check for any errors
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Get the input from the scanner
		input := scanner.Text()

		// Trim any leading/trailing spaces and newline characters
		input = strings.TrimSpace(input)

		// Split input into command name and arguments
		parts := strings.Fields(input)
		if len(parts) < 1 {
			fmt.Println("Please enter a command")
			continue
		}

		cmdName := parts[0]
		args := parts[1:]

		// Call the command if registered
		if cmdFnc, exists := commands.Commands[cmdName]; exists {
			cmdFnc(args) // Pass the arguments to the command function
		} else {
			fmt.Println("Command not found:", cmdName)
		}
	}
}
