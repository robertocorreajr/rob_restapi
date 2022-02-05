package repository

import (
	"errors"
	"rob_restapi/entity"
)

type BookRepositoryMemory struct {
	Books []entity.Book
	// db BookMemoryDb
}

func NewBookRepositoryMemory() entity.BookRepository {
	//TODO não receber nada na execução da função e criar o slice dentro da função.
	return &BookRepositoryMemory{Books: make([]entity.Book, 0)}
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
