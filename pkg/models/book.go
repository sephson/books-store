package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sephson/books-store/bookstore-api/pkg/config"
	"github.com/sephson/books-store/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.model
	Name string `gorm:**json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(db)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id, int64) {*Book, *gotm.DB}{
	var getBook Book
	db:=db.where("ID=7", ID).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.where("ID=7", ID).Delete(book)
	return book
}