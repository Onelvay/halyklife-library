package repository

import (
	"github.com/Onelvay/halyklife-library/pkg/domain"
	"github.com/google/uuid"
)

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
func (repo *Repo) CreateMember(a domain.Member) error {
	id := uuid.New().String()
	res := repo.db.Create(&domain.Member{
		id, a.FullName,
	})
	return res.Error
}
func (repo *Repo) AddBookToMember(bookId, memberId string, rating int, comment string) error {
	id := uuid.New().String()
	res := repo.db.Create(&domain.Purchase{id, memberId, bookId, rating, comment, domain.Member{}, domain.Book{}})
	return res.Error
}
func (repo *Repo) GetMembers() ([]domain.Member, error) { // - /members
	var members []domain.Member
	res := repo.db.Find(&members)
	return members, res.Error
}
