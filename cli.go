package main

import (
	"bufio"
	"fmt"
	"strings"
)

// GetInput prompts the user and reads their input
func GetInput(prompt string, reader *bufio.Reader) string {
	fmt.Println(prompt)
	var input string
	var err error

	input, err = reader.ReadString('\n')

	if err != nil {
		fmt.Println("\n Error reading input:", err)
		return ""
	}

	return strings.TrimSpace(input)
}
