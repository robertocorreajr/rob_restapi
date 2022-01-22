package main

import (
	"fmt"
	"rob_restapi/entity"
	"rob_restapi/repository"
	"rob_restapi/service"
)

// func (b Book) Add() uint64 {
// 	return 1
// }

// // Init books var as slice Book struct
// var books = make([]Book, 0)
// var newId = 3
// var mu sync.Mutex

// // Get All Books
// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	mu.Lock()
// 	json.NewEncoder(w).Encode(books)
// 	mu.Unlock()
// }

// // Get Single Book
// func getBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) // Get params

// 	// Loop through books and find with id
// 	mu.Lock()
// 	defer mu.Unlock()
// 	for _, book := range books {
// 		if book.ID == params["id"] {
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}
// 	// json.NewEncoder(w).Encode(&Book{})

// }

// // Create a New Book
// func createBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	mu.Lock()
// 	var book Book
// 	_ = json.NewDecoder(r.Body).Decode(&book)

// 	book.ID = strconv.Itoa(newId)
// 	books = append(books, book)
// 	newId++
// 	json.NewEncoder(w).Encode(book)
// 	mu.Unlock()
// }

// // Update a Book
// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)

// 	mu.Lock()
// 	defer mu.Unlock()
// 	for index, book := range books {
// 		if book.ID == params["id"] {
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			books[index] = book
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}
// 	// json.NewEncoder(w).Encode(books)
// }

// // Delete a Book
// func deleteBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)

// 	mu.Lock()
// 	defer mu.Unlock()
// 	for index, book := range books {
// 		if book.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			fmt.Println(books)
// 			json.NewEncoder(w).Encode(book)
// 			break
// 		}
// 	}
// 	// json.NewEncoder(w).Encode(books)
// }

func main() {

	db := repository.BookMemoryDb{Books: []entity.Book{}}
	repositoryMemory := repository.NewBookRepositoryMemory(db)

	service := service.NewBookService(repositoryMemory)

	bok, _ := service.Add("Titulo")
	fmt.Println(bok)
	// Init Router
	// r := mux.NewRouter()

	// // Mock Data - @todo - implement DB
	// books = append(books, Book{ID: "1", Isbn: "448743", Title: "Goland Book", Author: Author{Firstname: "Jhon", Lastname: "Doe"}})
	// books = append(books, Book{ID: "2", Isbn: "654942", Title: "Book Two", Author: Author{Firstname: "Steve", Lastname: "Smith"}})

	// // Route Handlers / Endpoints / Concurrency Goroutines - This works here? I down't know =D
	// go r.HandleFunc("/api/books", getBooks).Methods("GET")
	// go r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	// go r.HandleFunc("/api/books", createBook).Methods("POST")
	// go r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// go r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// log.Fatal(http.ListenAndServe(":8000", r))
}
