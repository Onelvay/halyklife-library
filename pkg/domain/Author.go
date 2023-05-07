package domain

// Автора характеризуют следующие параметры:
// - Идентификатор
// - ФИО
// - Псевдоним
// - Специализация
type Author struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	Specialization string `json:"specialization"`
	FullName       string `json:"full_name"`
}
