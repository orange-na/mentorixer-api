# Get started

## Prerequisite

- Go 1.22.2
- Docker (for running MySQL)

## Installation

1. Clone the repository:

   ```
   $ git clone https://github.com/orange-na/mentorixer-api.git
   $ cd mentorixer-api
   ```

2. Install dependencies:
   ```
   $ cp .env.sample .env
   $ go mod download
   ```

## Running the Application

1. Start the MySQL database using Docker Compose:

   ```
   $ docker-compose up -d
   ```

2. Run the application:
   ```
   $ go run main.go
   ```

## Building the Application

```
$ go build -o mentorixer-api
```

# Project structure

```
.
├── .env
├── .env.sample
├── .gitignore
├── cmd
│   └── server
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── handler
│   ├── model
│   └── router
├── jobs
├── main.go
├── middleware
│   ├── auth.go
│   └── cors.go
├── pkg
│   ├── db
│   └── external
└── utils
    └── token.go
```
