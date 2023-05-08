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

func (repo *Repo) GetMemberBooks(memberId string) ([]domain.Book, error) { // список книг взятых читателем
	//select * from books where books.id in (select p.book_id from members inner join purchases p on  member_id= %s and members.id = p.member_id)
	var books []domain.Book
	var bookIds []string
	repo.db.Table("members").
		Select("p.book_id").
		Joins("inner join purchases p on member_id = ? and members.id = p.member_id", memberId).
		Scan(&bookIds)
	repo.db.Find(&books, bookIds)
	return books, nil
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

func (repo *Repo) GetAuthors() ([]domain.Author, error) { // - /authors
	var authors []domain.Author
	res := repo.db.Find(&authors)
	fmt.Println(authors)
	return authors, res.Error
}

func (repo *Repo) GetBooks() ([]domain.Book, error) { // - /books
	var books []domain.Book
	res := repo.db.Find(&books)
	return books, res.Error
}

func (repo *Repo) GetMembers() ([]domain.Member, error) { // - /members
	var members []domain.Member
	res := repo.db.Find(&members)
	return members, res.Error
}
