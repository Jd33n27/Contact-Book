package main

import (
	"regexp"
	"strings"
)

// IsValidEmail validates if the provided email string is in a valid email format
// It uses a regex pattern to check for: localpart@domain.extension
func IsValidEmail(email string) bool {
	// Regular expression pattern for email validation
	// ^[a-zA-Z0-9._%+-]+ - starts with alphanumeric, dots, underscores, percent, plus, or hyphen
	// @ - must have an @ symbol
	// [a-zA-Z0-9.-]+ - domain name with alphanumeric, dots, or hyphens
	// \.[a-zA-Z]{2,}$ - ends with a dot followed by at least 2 letters (like .com, .org, etc)
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(pattern, email)
	if err != nil {
		return false
	}
	return match
}

// IsValidPhone validates if the provided phone string contains only digits
// It strips common phone formatting characters and checks if only numbers remain
func IsValidPhone(phone string) bool {
	// Remove common phone formatting characters: spaces, hyphens, parentheses, plus signs
	cleaned := strings.NewReplacer(
		" ", "",
		"-", "",
		"(", "",
		")", "",
		"+", "",
	).Replace(phone)

	// Regular expression to check if the cleaned string contains ONLY digits
	// ^[0-9]+$ means: start of string, one or more digits, end of string
	pattern := `^[0-9]+$`
	match, err := regexp.MatchString(pattern, cleaned)
	if err != nil {
		return false
	}

	// Also check that the phone number has a reasonable length (7-15 digits)
	if len(cleaned) < 7 || len(cleaned) > 15 {
		return false
	}

	return match
}

// IsValidName validates if the provided name string is not empty
// It allows letters, numbers, spaces, special characters, and emojis
func IsValidName(name string) bool {
	// Trim whitespace from the name
	name = strings.TrimSpace(name)

	// Check if name is empty
	if name == "" {
		return false
	}

	// Regular expression to check if name contains valid characters
	// Allows letters, numbers, spaces, special characters, and emojis
	// It basically allows everything except control characters
	pattern := `^[^\x00-\x1F\x7F]+$`
	match, err := regexp.MatchString(pattern, name)
	if err != nil {
		return false
	}

	return match
}
