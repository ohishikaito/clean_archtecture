package domain

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
}

type Users []User
