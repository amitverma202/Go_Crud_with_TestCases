# Library Management API

This API provides endpoints to manage a library system, allowing users to perform CRUD operations on books.

## Setup

### 1. Clone the Repository

git clone https://github.com/amitverma202/Go-CRUD-Project.git

CREATE DATABASE readershaven;

### 2. Navigate to the Project Directory

cd Simple-Library-Project

### 3. Install Dependencies

go mod tidy

### 4. PostgreSQL Database Setup

### a. Install PostgreSQL
If you haven't installed PostgreSQL, download and install it from here.

### b. Create Database
Create a new database named readershaven:
createdb readershaven

### c. Create Books Table
Connect to the readershaven database and create a books table with the following schema:

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    author VARCHAR(255),
    published_date DATE
);

### 5. Connect PostgreSQL to Go Application
Update the database connection string in main.go:

Replace the line:
db, err = sql.Open("postgres", "postgres://postgres:root@localhost/readershaven?sslmode=disable")
with:
db, err = sql.Open("postgres", "postgres://postgres:root@localhost/readershaven?sslmode=disable")

### 6. Run the Application

go run main.go
The server will start running at http://localhost:8080.

Base URL
The base URL for accessing the API is http://localhost:8080.

### 7. Instructions for Run Test suite

## for main
CREATE DATABASE readershaven;

## For test cases
CREATE DATABASE readershaven_test;

## Install Dependency
go mod tidy

## Run the application:
go run main.go

## Run Tests
To run the tests, execute the following command:

go test -v

This will run all the unit tests, integration tests, and API endpoint tests to ensure that the API functions correctly under various scenarios.

## Endpoints

/books: GET (Get all books)
/books/add: POST (Add a new book)
/books/update/{id}: PUT (Update a book by ID)
/books/delete/{id}: DELETE (Delete a book by ID)

## Status Codes

200 OK: Successful request
201 Created: Resource created
400 Bad Request: Invalid request data or parameters
404 Not Found: Resource not found
500 Internal Server Error: Server error

## Error Responses
The API will respond with appropriate HTTP status codes and error messages in case of any errors.

## Rate Limiting
There's no rate limiting implemented in this API.

## Conclusion

In this `README.md`:

- Added steps to set up PostgreSQL and create the `books` table.
- Updated the database connection string instruction in the setup steps.
- Included endpoints and status codes information.

Make sure to update the connection string and any other placeholders with your actual details. This `README.md` provides a comprehensive guide to setting up and running your Go API with PostgreSQL.