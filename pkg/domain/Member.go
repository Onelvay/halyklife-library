package domain

// Читатель:
// - Идентификатор
// - ФИО
// - Список взятых книг
type (
	Member struct {
		Id       string `json:"id"`
		FullName string `json:"full_name"`
	}
	Purchase struct {
		MemberId string
		BookId   string
		Member   Member `gorm:"references:Id"`
		Book     Book   `gorm:"references:Id"`
	}
)
