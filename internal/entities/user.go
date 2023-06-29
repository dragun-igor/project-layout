package entities

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Login    string    `json:"login" db:"login"`
	Password string    `json:"-" db:"password"`
	Name     string    `json:"name" db:"name"`
	//...
}
