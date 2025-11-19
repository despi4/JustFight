package models

// User - структура для представления пользователя
type User struct {
	ID       uint    `json:"id" db:"id"`             // Уникальный ID пользователя
	Username string  `json:"username" db:"username"` // Имя пользователя
	Email    string  `json:"email" db:"email"`       // Email пользователя
	Password string  `json:"-" db:"password"`        // Хешированный пароль (не возвращается в ответах)
	Wins     int     `json:"wins" db:"wins"`         // Количество побед
	Losses   int     `json:"losses" db:"losses"`     // Количество поражений
	Matches  []Match `json:"matches" db:"-"`         // Список матчей (связь с матчами)
}


