package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book 

// Get All Books 

func getBooks(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find with id
	for _, item := range books {
		if  item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

// Create a new Book
func createBook(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	var book Book 
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000)) // Mock ID - not safe
	books = append(books,book)
	json.NewEncoder(w).Encode(book)

}

// update a Book
func updateBook(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index],books[index+1:]...)
			var book Book 
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books,book)
			json.NewEncoder(w).Encode(book)
			return 
		}
	}	
	json.NewEncoder(w).Encode(books)
}

// Delete a Book
func deleteBook(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index],books[index+1:]...)
			break
		}
	}	
	json.NewEncoder(w).Encode(books)
}


// Mock Data -@todo - Implement DB
func main()  {
	books = append(books,Book{ID: "1",Isbn: "4321234",Title: "Book One",Author: &Author {Firstname: "John",Lastname: "Doe"}})
	books = append(books,Book{ID: "2",Isbn: "4322314",Title: "Book Two",Author: &Author {Firstname: "Aldy",Lastname: "Rifaldi.B"}})
	// // Init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/book",createBook).Methods("POST")
	r.HandleFunc("/api/book/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/api/book/{id}",deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",r))
}