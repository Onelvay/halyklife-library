package domain

// Книга:
// - Иденификатор
// - Название
// - Жанр
// - Код ISBN
type Book struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	ISBN     string `json:"ISBN"`
	AuthorId string `json:"author_id"`
	Author   Author `gorm:"references:Id"`
}
