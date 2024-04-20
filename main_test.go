package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


func setupTestDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost/readershaven_test?sslmode=disable")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS books (id SERIAL PRIMARY KEY, title VARCHAR(255), author VARCHAR(255), published_date DATE)")
	if err != nil {
		panic(err)
	}
	return db
}

func TestAddBook(t *testing.T) {
	db = setupTestDB()

	newBook := Book{
		Title:         "Test Book",
		Author:        "Test Author",
		PublishedDate: time.Now(),
	}
	body, _ := json.Marshal(newBook)
	req, err := http.NewRequest("POST", "/books/add", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addBook)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var responseBook Book
	json.Unmarshal(rr.Body.Bytes(), &responseBook)
	assert.Equal(t, newBook.Title, responseBook.Title)
	assert.Equal(t, newBook.Author, responseBook.Author)
	assert.Equal(t, newBook.PublishedDate.Format("2006-01-02"), responseBook.PublishedDate.Format("2006-01-02"))
}

func TestCRUDOperations(t *testing.T) {
	db = setupTestDB()

	newBook := Book{
		Title:         "Test Book",
		Author:        "Test Author",
		PublishedDate: time.Now(),
	}
	body, _ := json.Marshal(newBook)
	req, err := http.NewRequest("POST", "/books/add", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addBook)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)

	var responseBook Book
	json.Unmarshal(rr.Body.Bytes(), &responseBook)

	newBook.Title = "Updated Test Book"
	body, _ = json.Marshal(newBook)
	req, err = http.NewRequest("PUT", fmt.Sprintf("/books/update/%d", responseBook.ID), bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(updateBook)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	req, err = http.NewRequest("DELETE", fmt.Sprintf("/books/delete/%d", responseBook.ID), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(deleteBook)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetAllBooks(t *testing.T) {
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllBooks)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var allBooks []Book
	json.Unmarshal(rr.Body.Bytes(), &allBooks)
	assert.NotEmpty(t, allBooks)
}
