package notesmodel

import (
	"gorm.io/gorm"
)

type Notes struct {
	gorm.Model
	NotesData string
	UsersID int
}
