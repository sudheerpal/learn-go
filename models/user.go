package models

import (
	"time"
)

type User struct {
	Id         uint      `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email" gorm:"type:varchar(100);unique"`
	Mobile     string    `json:"mobile"`
	Department string    `json:"department"`
	Role       string    `json:"role"`
	Password   []byte    `json:"-"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}
