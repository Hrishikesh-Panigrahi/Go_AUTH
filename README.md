# RESTful-API-in-Go
A Authentication built with Go, GORM, and the Gin framework.

## Features

- User can Registration
- User Login
- Middleware

## Technologies Used
- Backend: Go, PostgreSQL
- Containerization: Docker

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

## Usage

1. **Run the application:**

```sh
   air
```

## Docker

1. **Pull the Docker Image:**
```sh
docker pull hrishikeshpanigrahi025/Go_AUTH
```

2. **Run the Docker container:**
```sh
docker run -p 8000:8000 hrishikeshpanigrahi025/Go_AUTH
```

## Project Structure
The project structure follows a standard Go project layout:

```
GoCMS/
├── controllers/
├── initializers/
├── middleware/
├── models/
├── Dockerfile
├── main.go
└── go.mod
```

- initializers/: Database initializers
- controllers/: Handlers for the different routes
- middleware/: Handling middleware   
- models/: Database models
- Dockerfile: Docker configuration
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
