package databaseConnect

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"gorm.io/driver/postgres"
	"../model/usermodel"
	"../model/notesmodel"
)

var DB *gorm.DB
var err error


func InitDatabase(){
	dsn := "host=127.0.0.1 user=shivam password=pass dbname=authen_app port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println("Successfully connected to database")
	}
	DB.AutoMigrate(&usermodel.Users{})
	DB.AutoMigrate(&notesmodel.Notes{})
}