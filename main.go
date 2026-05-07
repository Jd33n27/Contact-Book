package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// main is the entry point of the contact manager application
// It displays a menu loop for the user to interact with contact operations
func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Contact Manager CLI ---")
		fmt.Println("1. Add Contact")
		fmt.Println("2. List Contacts")
		fmt.Println("3. Edit Contact")
		fmt.Println("4. Delete Contact")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var option string
		var err error
		option, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		option = strings.TrimSpace(option)

		switch option {
		case "1":
			AddContact(reader)
		case "2":
			ListContacts()
		case "3":
			EditContact(reader)
		case "4":
			DeleteContact(reader)
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
