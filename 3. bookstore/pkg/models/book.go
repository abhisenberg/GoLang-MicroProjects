package models

/**
Model layer contains the model struct and the operations related to the model in the DB
**/

import (
	"github.com/abhisenberg/bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

//An instance of DB
var db *gorm.DB

type Book struct {
	gorm.Model         //This embeds the default gorm model fields into this struct, like ID, createdAt etc.
	Name        string `gorm:""json:"name"` //❗❗ What is this statement in backticks?
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}) //This statement informs the database about the new table (or changed table) and performs the necessary migrations.
}

func (b *Book) CreateBook() *Book { //❗❗ Why is this function an extension for book pointer?
	db.NewRecord(b) //It returns true or false as to whether the object we're passing to it already exists in the DB (by comparing the primary key)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books) //❗❗ What's the syntax of this Find function? What other parameter does it take?
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook) //❗❗ What's this syntax?
	return &getBook, db                       //❗❗ Why returning DB?
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book) //❗❗ What's this syntax?
	return book
}
