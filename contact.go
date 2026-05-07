package main

// Contact represents a contact with ID, FirstName, Surname, Phone, and Email fields
type Contact struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	Surname   string `json:"surname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

// ContactFile is the name of the JSON file where contacts are stored
const ContactFile string = "contacts.json"
