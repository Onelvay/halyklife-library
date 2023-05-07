package domain

// Автора характеризуют следующие параметры:
// - Идентификатор
// - ФИО
// - Псевдоним
// - Специализация
type Author struct {
	Id             string `json:"id" gorm:"primary_key"`
	Username       string `json:"username"`
	Specialization string `json:"specialization"`
	FullName       string `json:"full_name"`
}
