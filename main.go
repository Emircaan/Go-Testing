package main

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {

	books = append(books, Book{ID: 1, Title: "Kitap 1", Author: "Yazar 1"})
	books = append(books, Book{ID: 2, Title: "Kitap 2", Author: "Yazar 2"})

	http.HandleFunc("/getAll", getAllBooks)

	http.ListenAndServe(":8080", nil)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
