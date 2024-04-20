package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type Book struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"published_date"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:root@localhost/readershaven?sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/books", getAllBooks)
	http.HandleFunc("/books/add", addBook)
	http.HandleFunc("/books/update/", updateBook)
	http.HandleFunc("/books/delete/", deleteBook)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, author, published_date FROM books")
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var allBooks []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedDate)
		if err != nil {
			http.Error(w, "Failed to scan book", http.StatusInternalServerError)
			return
		}
		allBooks = append(allBooks, book)
	}

	jsonData, err := json.Marshal(allBooks)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO books (title, author, published_date) VALUES ($1, $2, $3)", newBook.Title, newBook.Author, newBook.PublishedDate)
	if err != nil {
		http.Error(w, "Failed to add book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid request URL", http.StatusBadRequest)
		return
	}
	bookID, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var updatedBook Book
	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE books SET title=$1, author=$2, published_date=$3 WHERE id=$4", updatedBook.Title, updatedBook.Author, updatedBook.PublishedDate, bookID)
	if err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedBook)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid request URL", http.StatusBadRequest)
		return
	}
	bookID, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM books WHERE id=$1", bookID)
	if err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Book with ID %d deleted successfully", bookID)
}
