package repository

import (
	"fmt"
	"github.com/Onelvay/halyklife-library/pkg/domain"
	"github.com/google/uuid"
)

func (repo *Repo) GetBooksByAuthor(authorId string) {
	var books []domain.Book
	repo.db.Where("author_id = ?", authorId).Find(&books)
	fmt.Println(books)
}

func (repo *Repo) CreateBook(authorId string, book domain.Book) error {
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
