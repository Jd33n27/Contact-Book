# Validation System Guide

## Overview

This document explains the validation system in the contact manager application. The validation logic has been separated into its own `validator.go` file, keeping the code organized and reusable.

---

## File: validator.go

This file contains three main validation functions that check user input before saving contacts.

### 1. IsValidEmail(email string) bool

**Purpose:** Validates if an email address is in proper format.

**How It Works:**

```go
pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
match, err := regexp.MatchString(pattern, email)
```

**Regex Breakdown:**

- `^` - Start of string
- `[a-zA-Z0-9._%+-]+` - Local part (before @): letters, numbers, dots, underscores, percent, plus, hyphen
- `@` - Required @ symbol
- `[a-zA-Z0-9.-]+` - Domain name: letters, numbers, dots, hyphens
- `\.` - Required dot (escaped with backslash)
- `[a-zA-Z]{2,}` - Extension: at least 2 letters (.com, .org, .co.uk, etc.)
- `$` - End of string

**Examples:**

- ✓ Valid: `john@example.com`, `user.name+tag@domain.co.uk`
- ✗ Invalid: `john@`, `@example.com`, `john.example.com`, `john@domain`

---

### 2. IsValidPhone(phone string) bool

**Purpose:** Validates that a phone number contains only digits (after stripping formatting characters).

**How It Works:**

**Step 1: Clean the phone number**

```go
cleaned := strings.NewReplacer(
    " ", "",
    "-", "",
    "(", "",
    ")", "",
    "+", "",
).Replace(phone)
```

This removes common formatting characters:

- Spaces: `123 456 7890` → `1234567890`
- Hyphens: `123-456-7890` → `1234567890`
- Parentheses: `(123) 456-7890` → `1234567890`
- Plus sign: `+1 234 567 8900` → `1234567890`

**Step 2: Validate with regex**

```go
pattern := `^[0-9]+$`
match, err := regexp.MatchString(pattern, cleaned)
```

Checks that the cleaned string contains ONLY digits (0-9).

**Step 3: Check length**

```go
if len(cleaned) < 7 || len(cleaned) > 15 {
    return false
}
```

Ensures the phone number is realistic (between 7 and 15 digits).

**Examples:**

- ✓ Valid: `1234567890`, `+1 (123) 456-7890`, `123 456 7890`
- ✗ Invalid: `123abc7890` (contains letters), `123456` (too short), `12345678901234567890` (too long)

---

### 3. IsValidName(name string) bool

**Purpose:** Validates that a name is not empty and contains valid characters including letters, numbers, spaces, hyphens, apostrophes, periods, emojis, and Unicode characters. It only disallows control characters.

**How It Works:**

**Step 1: Trim whitespace**

```go
name = strings.TrimSpace(name)
```

Removes leading and trailing spaces.

**Step 2: Check if empty**

```go
if name == "" {
    return false
}
```

Rejects empty names.

**Step 3: Validate with regex**

```go
pattern := `^[^\x00-\x1F\x7F]+$`
match, err := regexp.MatchString(pattern, name)
```

Checks that name does NOT contain control characters or DEL character:

**Allowed Characters:**

- Lowercase letters (a-z)
- Uppercase letters (A-Z)
- Numbers (0-9)
- Spaces
- Hyphens (-)
- Apostrophes (')
- Periods (.)
- Emojis (😀, 🎉, 👍, etc.)
- Unicode characters from any language (中文, العربية, ελληνικά, etc.)

**Disallowed Characters:**

- Control characters (\x00-\x1F)
- DEL character (\x7F)

**Examples:**

- ✓ Valid: `John`, `Mary Jane`, `Jean-Pierre`, `O'Brien`, `J.R.R. Tolkien`, `Anne-Marie123`
- ✓ Valid: `José García`, `François`, `李明`, `محمد`, `Παναγιώτης`, `Владимир`
- ✓ Valid: `John 😊`, `Alex 🎉`, `Sam 👍`, `Anna 🌟`
- ✗ Invalid: `` (empty), strings with control characters

---

## How Validators Are Used

### In AddContact Function

```go
func AddContact(reader *bufio.Reader) {
    // Get first name input
    var firstName string = GetInput("Enter First Name: ", reader)

    // Validate it
    if !IsValidName(firstName) {
        fmt.Println("First name must not be empty and contain only valid characters.")
        return  // Exit early if invalid
    }

    // Similar validation for surname and phone...

    // Validate email only if user provided one
    var email string = GetInput("Enter Email (optional): ", reader)
    if email != "" && !IsValidEmail(email) {
        fmt.Println("Email format is invalid. Please enter a valid email address.")
        return
    }

    // Only create contact if all validations pass
    var newContact Contact = Contact{
        ID:        newID,
        FirstName: firstName,
        Surname:   surname,
        Phone:     phone,
        Email:     email,
    }

    // Save to file
    SaveContacts(contacts)
}
```

### In EditContact Function

The same validation functions are called when editing each field:

```go
var newPhone string = GetInput(fmt.Sprintf("New Phone [%s]: ", c.Phone), reader)
if newPhone != "" {
    if !IsValidPhone(newPhone) {
        fmt.Println("Phone number is invalid. Please enter only numbers (7-15 digits).")
        return
    }
    contacts[i].Phone = newPhone
}
```

---

## Contact Struct Changes

```go
type Contact struct {
    ID        string
    FirstName string
    Surname   string
    Phone     string
    Email     string
}
```

## Data Flow Example

When a user adds a contact:

1. **User Input** → `GetInput()` prompts user and reads input
2. **Validation** → `IsValidName()`, `IsValidPhone()`, `IsValidEmail()` check input
3. **Decision**:
   - If invalid → Show error message, return early
   - If valid → Continue to next step
4. **Create Contact** → Build Contact struct with validated data
5. **Save** → `SaveContacts()` writes to JSON file

---

## Testing the Validators

Try these inputs to test:

**Email Validation:**

```
Valid:   john@example.com
Invalid: john.example.com (no @)
Invalid: john@domain (no extension)
```

**Phone Validation:**

```
Valid:   (123) 456-7890
Valid:   +1 123 456 7890
Invalid: 123-ABC-7890 (contains letters)
Invalid: 12345 (too short)
```

**Name Validation:**

```
Valid:   John
Valid:   Jean-Pierre
Valid:   O'Brien
Valid:   J.R.R. Tolkien
Valid:   Anne-Marie123
Valid:   José García
Valid:   李明 (Chinese)
Valid:   محمد (Arabic)
Valid:   John 😊 (with emoji)
Valid:   Alex 🎉
Invalid: (empty)
Invalid: Strings with control characters
```

---

## Key Benefits of This Architecture

1. **Separation of Concerns** - Validation logic is isolated in its own file
2. **Reusability** - Validator functions can be used by multiple operations (Add, Edit, Delete)
3. **Maintainability** - Easy to update validation rules in one place
4. **Testability** - Validators can be tested independently
5. **Clarity** - Each function has a single, clear responsibility
