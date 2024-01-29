package model

import (
	"github.com/jinzhu/gorm"
	"github.com/xiayulin123/GO_GameStoreAPI/pkg/config"
)

var db *gorm.DB

type Toys struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Price       string `gorm:"column:price" json:"price"`
	Description string `gorm:"column:description" json:"description"`
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

func GetToyById(Id int64) (*Toys, *gorm.DB) {
	var Toy Toys
	db := db.Where("ID=?", Id).Find(&Toy)
	return &Toy, db
}
func DeleteToy(Id int64) Toys {
	var Toy Toys
	db.Where("ID=?", Id).Delete(&Toy)
	return Toy
}
