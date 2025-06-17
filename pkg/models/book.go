package models 

import (
	"gorm.io/gorm"
	"github.com/abh1SHAKE/go-crud/pkg/config"
)

var db *gorm.DB 

type Book struct {
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}) // Automatically creates the table if it doesn't exist
}