# Contact Book Project Overview

A comprehensive guide explaining the What, Why, Who, Where, When, and How of the Contact Manager application.

---

## What?

**The Contact Book** is a command-line contact management application built in Go that allows users to store, organize, and manage personal contact information.

### Core Functionality:

- **Add Contacts** - Create new contacts with validated first name, surname, phone number, and email
- **List Contacts** - View all saved contacts in a formatted display
- **Edit Contacts** - Search and update existing contact information with flexible search options
- **Delete Contacts** - Search and remove contacts from the database
- **Data Persistence** - All contacts are automatically saved to a JSON file (`contacts.json`)

### Key Features:

- Input validation for emails, phone numbers, and names (including emoji support)
- Flexible search functionality (search by First Name, Surname, Phone, or Email)
- Persistent storage using JSON format
- Modular architecture with separated concerns
- Command-line interface with interactive menu

---

## Why?

The Contact Manager was built to demonstrate several software engineering principles:

### 1. **Code Organization & Modularity**

- Separate concerns into different files (contact.go, storage.go, operations.go, cli.go, validator.go)
- Each file has a single, clear responsibility
- Makes code easier to maintain, test, and extend

### 2. **Data Validation & Integrity**

- Ensures only valid data is saved to the database
- Prevents corrupted or malformed entries
- Provides clear error messages to guide users

### 3. **Real-World Application Development**

- Demonstrates CRUD operations (Create, Read, Update, Delete)
- Shows how to work with file I/O and JSON serialization
- Practical command-line interface design

### 4. **Best Practices in Go**

- Proper package structure
- Verbose, readable code with clear variable declarations
- Comprehensive comments and documentation
- Error handling patterns

---

## Who?

### Primary Users:

- **End Users** - Anyone needing a simple, command-line contact management tool
- **Developers/Students** - Learning Go programming and software architecture
- **Developers** - Using this as a reference for modular Go application structure

### Stakeholders:

- **Project Creator** - Built as a learning/practice project
- **Code Reviewers** - Evaluating code quality and architectural decisions
- **Future Contributors** - Developers who might extend or maintain the application

---

## Where?

### Execution Environment:

- **Operating System** - Linux/Unix-based systems (or any OS with Go installed)
- **Runtime** - Go version 1.26+ (or compatible)
- **File System** - Application directory where `contacts.json` is stored
- **Terminal/Console** - Standard command-line interface

### File Structure:

```
contact_book/
├── main.go              # Entry point, menu loop
├── contact.go           # Contact struct definition
├── storage.go           # File I/O operations (LoadContacts, SaveContacts)
├── operations.go        # CRUD operations (Add, List, Edit, Delete)
├── cli.go              # User input handling
├── validator.go        # Input validation logic
├── contacts.json       # Data storage (generated on first use)
├── go.mod              # Go module definition
└── README files        # Documentation
```

---

## When?

### Use Cases & Timing:

1. **Daily Use**
   - Managing personal or business contacts
   - Quick contact lookups
   - Adding new contact information

2. **Development/Testing**
   - Running during active development: `go run .`
   - Building executable: `go build`
   - Testing specific features

3. **Learning Scenarios**
   - Studying Go language fundamentals
   - Understanding software architecture patterns
   - Learning about data validation and persistence

### Workflow Timeline:

```
Start Application
    ↓
Display Menu
    ↓
User Selects Operation (1-5)
    ↓
Operation Executed with Validation
    ↓
Result Displayed
    ↓
Loop Back to Menu
    ↓
User Selects Exit (5)
    ↓
Application Terminates
```

---

## How?

### Architecture Overview:

The application uses a **layered architecture** pattern:

```
┌──────────────────────────────┐
│   Main (CLI Menu Loop)       │  - Presents options
│       main.go                │  - Routes user input
└──────────────┬───────────────┘
               │
┌──────────────▼──────────────────────┐
│  Operations Layer                   │  - Business Logic
│  (Add, List, Edit, Delete)          │
│  operations.go                      │
└──────────────┬──────────────────────┘
               │
    ┌──────────┴──────────┬──────────────┐
    │                     │              │
┌───▼────────┐  ┌────────▼────────┐  ┌─▼───────────────┐
│ Validators │  │ CLI Input       │  │ Storage          │
│ validator. │  │ cli.go          │  │ storage.go       │
│ go         │  │ GetInput()      │  │ Load/Save        │
└────────────┘  └─────────────────┘  └──────────────────┘
                          │
                ┌─────────▼──────────┐
                │  Data Layer        │
                │  contact.go        │
                │  Contact struct    │
                │  contacts.json     │
                └────────────────────┘
```

### Data Flow Example - Adding a Contact:

```
1. User selects "1. Add Contact"
2. GetInput() prompts for First Name
3. IsValidName() validates the input
4. If invalid → Show error, return to menu
5. If valid → Continue to Surname
6. Repeat validation for Surname, Phone, Email
7. Create Contact struct with validated data
8. LoadContacts() reads existing contacts from JSON
9. Append new contact to the list
10. SaveContacts() writes all contacts back to JSON
11. Display success message
12. Return to menu
```

### Technology Stack:

| Component    | Technology  | Purpose                             |
| ------------ | ----------- | ----------------------------------- |
| Language     | Go          | Compiled, fast, cross-platform      |
| Data Format  | JSON        | Human-readable, easy to parse       |
| Storage      | File System | Simple, no database required        |
| Interface    | CLI         | Command-line menu interface         |
| Input/Output | bufio, fmt  | Reading user input, printing output |
| Validation   | regexp      | Pattern matching for validation     |

### Key Implementation Details:

**1. Modular Package Structure**

- All files in `package main`
- Functions are exported (capitalized) for internal access
- No external dependencies required

**2. Validation Layer**

- `IsValidEmail()` - Uses regex to validate email format
- `IsValidPhone()` - Strips formatting, checks digits only (7-15 length)
- `IsValidName()` - Allows letters, numbers, spaces, hyphens, apostrophes, periods, and emojis

**3. Flexible Search**

- Add, Edit, Delete all support searching by multiple fields
- User chooses search criteria from menu
- Case-sensitive exact matching

**4. Data Persistence**

- Contacts stored in JSON format
- Automatic file creation on first use
- Pretty-printed with indentation for readability
- Loaded into memory for operations, written back after changes

---

## System Requirements

| Requirement          | Details                     |
| -------------------- | --------------------------- |
| **Go Version**       | 1.26+                       |
| **Operating System** | Linux, macOS, Windows       |
| **Disk Space**       | Minimal (~1MB)              |
| **Memory**           | <10MB                       |
| **Terminal**         | Any standard terminal/shell |

---

## Quick Start

```bash
# Build the application
go build

# Run the application
go run .

# Or run the compiled binary
./contact_book
```

Then interact with the menu by entering numbers 1-5 to manage your contacts.

---

<!--## Future Enhancements

Potential improvements to the application:

1. **Search Functionality** - Search contacts by partial name or phone number
2. **Categories** - Organize contacts by groups (friends, family, work)
3. **Export** - Export contacts to CSV or vCard format
4. **Import** - Import contacts from other formats
5. **Database** - Replace JSON with a database system
6. **Web Interface** - Create a REST API and web UI
7. **Sync** - Cloud synchronization across devices
8. **Backup** - Automatic backup and recovery features-->

