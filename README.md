# Go-AUTH

An Authentication API built with Go, GORM, and the Gin framework.

## Features

- API endpoint for User Registration
- API endpoint for User Login
- Middleware using JWT tokens

## Technologies Used

- Backend: Go, PostgreSQL

## Requirements

- Go 1.22+
- A SQL database (e.g., PostgreSQL, MySQL, SQLite)
- Git

## Installation

1. **Clone the repository:**

```sh
   git clone https://github.com/Hrishikesh-Panigrahi/Go_AUTH.git
   cd Auth-in-Go
```

2. **Install dependencies:**

```sh
   go mod tidy
```

And you are all set

## Project Structure

The project structure follows a standard Go project layout:

```
GoCMS/
├── controllers/
├── initializers/
├── middleware/
├── models/
├── main.go
└── go.mod
```

- initializers/: Database initializers
- controllers/: Handlers for the different routes
- middleware/: Handling middleware
- models/: Database models
- main.go: Entry point of the application
- go.mod: Go module file

## Run Locally

To run the project locally, you have 3 options:

### 1. Launch Debugger

- Open your project in Visual Studio Code.
- Set breakpoints as needed.
- Launch the debugger by pressing F5 or by selecting Run > Start Debugging from the menu.

### 2. Run Air:

Ensure you have Air installed for live reloading.

- Start Air by running the following command in your terminal:

```sh
   air
```

### 3. Run go run main.go Command

- Open your terminal.

- Navigate to the project directory.

- Run the following command to start the application:

```sh
   go run main.go
```
