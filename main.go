package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

//Book struct(model) kind of like a class
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//init books variable as a slice book struct, a slice is a variable length array
var books []Book

//get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

//get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get any params
	//looop through books and find id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//create book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//update books
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params :=mux.Vars(r)
	for index, item :=range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
		    var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
		
	}
	json.NewEncoder(w).Encode(books)
}

//delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params :=mux.Vars(r)
	for index, item :=range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
		    break
		}
		
	}
	json.NewEncoder(w).Encode(books)

}

func main() {
	//starting router
	r := mux.NewRouter()

	//sample data @todo implement db
	books = append(books, Book{ID: "1", Isbn: "123456", Title: "book one", Author: &Author{Firstname: "john", Lastname: "doe"}})
	books = append(books, Book{ID: "2", Isbn: "113456", Title: "book two", Author: &Author{Firstname: "bob", Lastname: "flow"}})

	//creating route handlers( endpoints)
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//running server
	log.Fatal(http.ListenAndServe(":8000", r))

}
