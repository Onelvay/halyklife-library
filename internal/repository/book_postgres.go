package repository

import (
	"github.com/Onelvay/halyklife-library/pkg/domain"
	"github.com/google/uuid"
)

func (repo *Repo) CreateBook(book domain.Book) error {
	id := uuid.New().String()
	res := repo.db.Create(&domain.Book{
		Id:       id,
		Name:     book.Name,
		Genre:    book.Genre,
		ISBN:     book.ISBN,
		AuthorId: book.AuthorId,
	})
	return res.Error
}

func (repo *Repo) UpdateBookNameById(id, name string) error {
	res := repo.db.Model(&domain.Book{}).Where("id = ?", id).Update("name", name)
	return res.Error
}

func (repo *Repo) DeleteBookById(id string) error {
	repo.db.Where("book_id = ?", id).Delete(&domain.Purchase{})
	res := repo.db.Where("id = ?", id).Delete(&domain.Book{})
	return res.Error
}

func (repo *Repo) GetBooks() ([]domain.Book, error) { // - /books
	var books []domain.Book
	res := repo.db.Find(&books)
	return books, res.Error
}
