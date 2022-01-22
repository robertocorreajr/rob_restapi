package entity

import uuid "github.com/satori/go.uuid"

type BookRepository interface {
	Get(id string) (Book, error)
	Add(book Book) (Book, error)
}

// Book Struct (Model)
type Book struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
	Author
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lasttname"`
}

func NewBook() *Book {
	book := Book{
		ID: uuid.NewV4().String(),
	}
	return &book
}
