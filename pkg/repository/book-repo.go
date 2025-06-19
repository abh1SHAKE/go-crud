package repository

import (
	"gorm.io/gorm"
	"github.com/abh1SHAKE/go-crud/pkg/config"
	"github.com/abh1SHAKE/go-crud/pkg/models"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&models.Book{})
}

func CreateBook(book *models.Book) error {
	return db.Create(book).Error
}

func GetAllBooks() ([]models.Book, error) {
	var books[] models.Book
	err := db.Find(&books).Error
	return books, err
}

func GetBookById(id int64) (*models.Book, error) {
	var book models.Book
	err := db.First(&book, id).Error
	return &book, err
}

func DeleteBook(id int64) (*models.Book, error) {
	var book models.Book
	err := db.First(&book, id).Error
	if err != nil {
		return nil, err
	}

	err = db.Delete(&book).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func UpdateBook(book *models.Book) error {
	return db.Save(book).Error
}