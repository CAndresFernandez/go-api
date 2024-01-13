package models

import (
	"github.com/jinzhu/gorm"
)


type Book struct {
	gorm.Model
	Name string `json:"name"`
	AuthorName string `json:"author_name"`
}