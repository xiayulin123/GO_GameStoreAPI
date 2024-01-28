package model

import (
	"github.com/jinzhu/gorm"
	"github.com/xiayulin123/Go-bookstore/pkg/config"
)

var db *gorm.DB

type Toys struct {
	gorm.Model
	Name        string  `gorm:"column:name" json:"name"`
	Price       float32 `gorm:"column:price" json:"price"`
	Description string  `gorm:"column:description" json:"description"`
}

func init() {
	config.Connection()
	db = config.GetDB()
	db.AutoMigrate(&Toys{})
}
func (b *Toys) CreateToy() *Toys {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetToys() []Toys {
	var Toy []Toys
	db.Find(&Toy)
	return Toy
}

func GetToyById(Id int64) (*Toys, error) {
	var Toy Toys
	result := db.Where("ID=?", Id).Find(&Toy)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Toy, nil
}
func DeleteToy(Id int64) Toys {
	var Toy Toys
	db.Where("ID=?", Id).Delete(&Toy)
	return Toy
}
