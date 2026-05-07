package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// SaveContacts writes the contacts slice to a JSON file
func SaveContacts(contacts []Contact) {
	var data []byte
	var err error

	data, err = json.MarshalIndent(contacts, " ", "\t")
	if err != nil {
		fmt.Println("Error saving contacts:", err)
		return
	}

	err = os.WriteFile(ContactFile, data, 0644)
	if err != nil {
		fmt.Println("Error Adding Contact:", err)
		return
	}
}

// LoadContacts reads contacts from the JSON file
func LoadContacts() ([]Contact, error) {
	var err error
	_, err = os.Stat(ContactFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Contact{}, nil
		}
		return nil, err
	}

	var contacts []Contact

	data, err := os.ReadFile(ContactFile)
	if err != nil {
		return contacts, err
	}

	err = json.Unmarshal(data, &contacts)
	if err != nil {
		fmt.Println("Error loading contacts:", err)
		return contacts, nil
	}

	return contacts, nil
}
