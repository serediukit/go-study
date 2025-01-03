package models

import (
	"github.com/jinzhu/gorm"
	"github.com/serediukit/go-study/3___go_my_sql/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}
