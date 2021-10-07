package usermodel

import (
	"gorm.io/gorm"
	"github.com/dgrijalva/jwt-go"
	"../notesmodel"
)

type Users struct{
	gorm.Model
	Username string
	Password string
	Notes []notesmodel.Notes `gorm:"foreignKey:UsersID"`
}

type Claims struct{
	Username string
	UsersID uint
	jwt.StandardClaims
}