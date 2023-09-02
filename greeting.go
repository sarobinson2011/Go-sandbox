package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a new reader to read input from the user
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user for their name
	fmt.Print("Enter your name: ")

	// Read the user's input as a string
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Print a personalized greeting
	fmt.Printf("Hello, %s! Welcome to the Go programming world!\n", name)
}
