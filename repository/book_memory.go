package repository

import (
	"errors"
	"rob_restapi/entity"
)

type BookMemoryDb struct {
	Books []entity.Book
}

type BookRepositoryMemory struct {
	db BookMemoryDb
}

func NewBookRepositoryMemory(db BookMemoryDb) *BookRepositoryMemory {
	return &BookRepositoryMemory{db: db}
}

func (c *BookRepositoryMemory) Get(id string) (entity.Book, error) {
	for _, book := range c.db.Books {
		if book.ID == id {
			return book, nil
		}
	}
	return entity.Book{}, errors.New("no book found with this ID")
}

func (c *BookRepositoryMemory) Add(book entity.Book) (entity.Book, error) {
	c.db.Books = append(c.db.Books, book)
	return book, nil
}
