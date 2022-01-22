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

func (b *BookService) Add(name string) (entity.Book, error) {
	book := entity.NewBook()
	book.Title = name
	bok, err := b.Repository.Add(*book)
	if err != nil {
		return entity.Book{}, err
	}
	return bok, nil
}
