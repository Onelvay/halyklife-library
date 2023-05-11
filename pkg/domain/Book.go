package domain

// Книга:
// - Иденификатор
// - Название
// - Жанр
// - Код ISBN
type Book struct {
	Id       string  `json:"id" gorm:"primary_key"`
	Name     string  `json:"name"`
	Genre    string  `json:"genre"`
	ISBN     string  `json:"ISBN"`
	Rating   float32 `json:"rating"`
	Price    float32 `json:"price"`
	AuthorId string  `json:"author_id"`
	Author   Author  `gorm:"references:Id" json:"-"`
}
