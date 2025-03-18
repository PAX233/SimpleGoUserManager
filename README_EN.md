# SimpleGoUserManager

A simple user management system implementation with basic CRUD operations using Go.

## Features

- Uses struct to store user information (name, age, gender, email)
- Implements data persistence with JSON file
- Supports adding new users
- Supports listing users and viewing individual details
- Supports updating user information
- Supports deleting users by ID

## Requirements

- Go 1.24 or higher

## Quick Start

1. Clone repository
```bash
git clone https://github.com/yourusername/SimpleGoUserManager.git
```

2. Enter project directory
```bash
cd basicSystem
```

3. Run program
```bash
go run main.go
```

## Project Structure

```
├── main.go         # Main entry point
├── customer/       # Customer management module
│   ├── add.go      # Add customer
│   ├── delete.go   # Delete customer
│   ├── update.go   # Update customer
│   ├── show.go     # Display customers
│   ├── types.go    # Data structures
│   ├── load.go     # Data loading
│   └── save.go     # Data saving
├── menu/           # Menu system
│   └── menu.go     # Menu rendering
├── data/           # Data storage
│   └── customers.json  # User data
├── go.mod          # Dependency management
└── README.md       # Documentation (Chinese)
```

## Sample Output

The program demonstrates:
- Loading user data from persistent storage
- Adding new users
- Listing users
- Updating user information
- Deleting users
- Automatic data saving

## License
MIT License

### Notes
- **Features**: Updated to reflect struct-based user storage
- **Structure**: Added `customer/` and `menu/` directories
- **Output**: Revised to accurately describe functionality