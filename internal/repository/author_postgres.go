package repository

import (
	"fmt"
	"github.com/Onelvay/halyklife-library/pkg/domain"
	"github.com/google/uuid"
)

func (repo *Repo) GetAuthorBooks(authorId string) ([]domain.Book, error) { // список книг автора которые есть в наличии
	var books []domain.Book
	//SELECT books FROM BOOKS INNER JOIN authors a on a.id=%s and books.author_id = a.id inner join purchases p on books.id != p.book_id
	res := repo.db.Table("books").Select("*").
		Joins("INNER JOIN authors a ON a.id = ? and books.author_id = a.id", authorId).
		Joins("inner join purchases p on books.id != p.book_id").
		Scan(&books)
	return books, res.Error
}

func (repo *Repo) CreateAuthor(a domain.Author) error {
	id := uuid.New().String()
	res := repo.db.Create(&domain.Author{
		id, a.Username, a.Specialization, a.FullName,
	})
	return res.Error
}

func (repo *Repo) GetAuthors() ([]domain.Author, error) { // - /authors
	var authors []domain.Author
	res := repo.db.Find(&authors)
	fmt.Println(authors)
	return authors, res.Error
}
