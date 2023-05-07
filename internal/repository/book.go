package repository

import (
	"github.com/Onelvay/halyklife-library/pkg/domain"
	"github.com/google/uuid"
)

// Специальные пути:
// - /authors/{id}/books - список книг автора которые есть в наличии
// - /members/{id}/books - список книг взятых читателем
func (repo *Repo) GetAuthorBooks(authorId string) ([]domain.Book, error) {
	var books []domain.Book
	//SELECT books FROM BOOKS INNER JOIN authors a on a.id=%s and books.author_id = a.id inner join purchases p on books.id != p.book_id
	res := repo.db.Table("books").Select("*").
		Joins("INNER JOIN authors a ON a.id = ? and books.author_id = a.id", authorId).
		Joins("inner join purchases p on books.id != p.book_id").
		Scan(&books)
	return books, res.Error
}

func (repo *Repo) GetMemberBooks() ([]domain.Book, error) {
	//select * from books where books.id in (select p.book_id from members inner join purchases p on  member_id= %s and members.id = p.member_id)
	var books []domain.Book
	res := repo.db.Table("books").Select("*").Where("books.id in ")

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
