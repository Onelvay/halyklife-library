package domain

// Читатель:
// - Идентификатор
// - ФИО
// - Список взятых книг
type (
	Member struct {
		Id       string `json:"id" gorm:"primary_key"`
		FullName string `json:"full_name"`
	}
	Purchase struct {
		Id       string `gorm:"primary_key"`
		MemberId string
		BookId   string
		Member   Member `gorm:"references:Id"`
		Book     Book   `gorm:"references:Id"`
	}
)
