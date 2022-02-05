package service

import "rob_restapi/entity"

type BookService struct {
	Repository entity.BookRepository
}

func NewBookService(repository entity.BookRepository) *BookService {
	return &BookService{Repository: repository}
}

func (b *BookService) findById(id string) (entity.Book, error) {
	book, error := b.Repository.Get(id)
	if error != nil {
		return entity.Book{}, error
	}
	return book, nil
}

func (b *BookService) Add(isbn, title, firstname, lastname string) (entity.Book, error) {
	book := entity.NewBook()
	book.Isbn = isbn
	book.Title = title
	book.Author.Firstname = firstname
	book.Author.Lastname = lastname
	bk, err := b.Repository.Add(*book)
	if err != nil {
		return entity.Book{}, err
	}
	return bk, nil
}
