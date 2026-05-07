package main

import (
	"bufio"
	"fmt"
	"time"
)

// AddContact creates a new contact and saves it to the contacts file
// It validates email format, phone number format, and names before saving
func AddContact(reader *bufio.Reader) {
	var firstName string = GetInput("Enter First Name: ", reader)
	if !IsValidName(firstName) {
		fmt.Println("First name must not be empty and contain only valid characters.")
		return
	}

	var surname string = GetInput("Enter Surname: ", reader)
	if !IsValidName(surname) {
		fmt.Println("Surname must not be empty and contain only valid characters.")
		return
	}

	var phone string = GetInput("Enter Phone: ", reader)
	if !IsValidPhone(phone) {
		fmt.Println("Phone number is invalid. Please enter only numbers (7-15 digits).")
		return
	}

	var email string = GetInput("Enter Email (optional): ", reader)
	if email != "" && !IsValidEmail(email) {
		fmt.Println("Email format is invalid. Please enter a valid email address.")
		return
	}

	var newID string = fmt.Sprintf("%d", time.Now().Unix())

	var newContact Contact = Contact{
		ID:        newID,
		FirstName: firstName,
		Surname:   surname,
		Phone:     phone,
		Email:     email,
	}

	contacts, err := LoadContacts()
	if err != nil {
		fmt.Println("Error loading contacts:", err)
		return
	}

	contacts = append(contacts, newContact)
	SaveContacts(contacts)

	fmt.Println("Contact added successfully!")
}

// ListContacts displays all contacts
func ListContacts() {
	contacts, err := LoadContacts()
	if err != nil {
		fmt.Println("Error loading contacts:", err)
		return
	}

	if len(contacts) == 0 {
		fmt.Println("No contacts found.")
		return
	}

	fmt.Println("\n--- Contact List ---")
	for _, c := range contacts {
		fmt.Printf("ID: %s\nFirst Name: %s\nSurname: %s\nPhone: %s\nEmail: %s\n\n", c.ID, c.FirstName, c.Surname, c.Phone, c.Email)
	}
}

// EditContact allows the user to update an existing contact
// It validates email and phone formats before saving changes
// User can search by FirstName, Surname, Phone, or Email
func EditContact(reader *bufio.Reader) {
	ListContacts()
	var contacts []Contact
	var err error

	contacts, err = LoadContacts()
	if err != nil {
		fmt.Println("Error loading contacts:", err)
		return
	}

	if len(contacts) == 0 {
		fmt.Println("No contacts found.")
		return
	}

	fmt.Println("\nSearch by:")
	fmt.Println("1. First Name")
	fmt.Println("2. Surname")
	fmt.Println("3. Phone Number")
	fmt.Println("4. Email")
	fmt.Print("Choose search field: ")

	var searchChoice string = GetInput("", reader)

	var searchTerm string
	var found bool = false
	var contactIndex int = -1

	switch searchChoice {
	case "1":
		searchTerm = GetInput("Enter First Name to edit: ", reader)
		for i, c := range contacts {
			if c.FirstName == searchTerm {
				found = true
				contactIndex = i
				break
			}
		}

	case "2":
		searchTerm = GetInput("Enter Surname to edit: ", reader)
		for i, c := range contacts {
			if c.Surname == searchTerm {
				found = true
				contactIndex = i
				break
			}
		}

	case "3":
		searchTerm = GetInput("Enter Phone Number to edit: ", reader)
		for i, c := range contacts {
			if c.Phone == searchTerm {
				found = true
				contactIndex = i
				break
			}
		}

	case "4":
		searchTerm = GetInput("Enter Email to edit: ", reader)
		for i, c := range contacts {
			if c.Email == searchTerm {
				found = true
				contactIndex = i
				break
			}
		}

	default:
		fmt.Println("Invalid choice.")
		return
	}

	if !found {
		fmt.Println("Contact not found.")
		return
	}

	i := contactIndex
	c := contacts[i]
	fmt.Printf("Editing '%s %s'. Leave blank to keep current value.\n", c.FirstName, c.Surname)

	var newFirstName string = GetInput(fmt.Sprintf("New First Name [%s]: ", c.FirstName), reader)
	if newFirstName != "" {
		if !IsValidName(newFirstName) {
			fmt.Println("First name must contain only valid characters.")
			return
		}
		contacts[i].FirstName = newFirstName
	}

	var newSurname string = GetInput(fmt.Sprintf("New Surname [%s]: ", c.Surname), reader)
	if newSurname != "" {
		if !IsValidName(newSurname) {
			fmt.Println("Surname must contain only valid characters.")
			return
		}
		contacts[i].Surname = newSurname
	}

	var newPhone string = GetInput(fmt.Sprintf("New Phone [%s]: ", c.Phone), reader)
	if newPhone != "" {
		if !IsValidPhone(newPhone) {
			fmt.Println("Phone number is invalid. Please enter only numbers (7-15 digits).")
			return
		}
		contacts[i].Phone = newPhone
	}

	var newEmail string = GetInput(fmt.Sprintf("New Email [%s]: ", c.Email), reader)
	if newEmail != "" {
		if !IsValidEmail(newEmail) {
			fmt.Println("Email format is invalid. Please enter a valid email address.")
			return
		}
		contacts[i].Email = newEmail
	}

	SaveContacts(contacts)

	fmt.Println("Contact updated successfully!")
}

// DeleteContact removes a contact by searching across multiple fields
// User can search by FirstName, Surname, Phone, or Email
func DeleteContact(reader *bufio.Reader) {
	ListContacts()
	var contacts []Contact
	var err error

	contacts, err = LoadContacts()
	if err != nil {
		fmt.Println("Error loading contacts:", err)
		return
	}

	if len(contacts) == 0 {
		fmt.Println("No contacts found.")
		return
	}

	fmt.Println("\nSearch by:")
	fmt.Println("1. First Name")
	fmt.Println("2. Surname")
	fmt.Println("3. Phone Number")
	fmt.Println("4. Email")
	fmt.Print("Choose search field: ")

	var searchChoice string = GetInput("", reader)

	var searchTerm string
	var found bool = false

	switch searchChoice {
	case "1":
		searchTerm = GetInput("Enter First Name to delete: ", reader)
		for i, c := range contacts {
			if c.FirstName == searchTerm {
				var fullName string = c.FirstName + " " + c.Surname
				contacts = append(contacts[:i], contacts[i+1:]...)
				SaveContacts(contacts)
				fmt.Printf("%s has been Deleted!\n", fullName)
				ListContacts()
				return
			}
		}

	case "2":
		searchTerm = GetInput("Enter Surname to delete: ", reader)
		for i, c := range contacts {
			if c.Surname == searchTerm {
				var fullName string = c.FirstName + " " + c.Surname
				contacts = append(contacts[:i], contacts[i+1:]...)
				SaveContacts(contacts)
				fmt.Printf("%s has been Deleted!\n", fullName)
				ListContacts()
				return
			}
		}

	case "3":
		searchTerm = GetInput("Enter Phone Number to delete: ", reader)
		for i, c := range contacts {
			if c.Phone == searchTerm {
				var fullName string = c.FirstName + " " + c.Surname
				contacts = append(contacts[:i], contacts[i+1:]...)
				SaveContacts(contacts)
				fmt.Printf("%s has been Deleted!\n", fullName)
				ListContacts()
				return
			}
		}

	case "4":
		searchTerm = GetInput("Enter Email to delete: ", reader)
		for i, c := range contacts {
			if c.Email == searchTerm {
				var fullName string = c.FirstName + " " + c.Surname
				contacts = append(contacts[:i], contacts[i+1:]...)
				SaveContacts(contacts)
				fmt.Printf("%s has been Deleted!\n", fullName)
				ListContacts()
				return
			}
		}

	default:
		fmt.Println("Invalid choice.")
		return
	}

	if !found {
		fmt.Println("Contact not found.")
	}
}
