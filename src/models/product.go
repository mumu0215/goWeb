package models

import (
	"fmt"
	db "src/database"
)

type Products struct {
	//gorm.Model
	ID      int    `gorm:"primary_key;AUTO_INCREMENT"`
	Product string `gorm:"not null"`
	Details string
}

func (p Products) TableName() string {
	return "products"
}

func (p *Products) AddProducts() {
	fmt.Println(p)
	db.MyDb.Create(p)
}

//must provide primary key
func (p *Products) DelProducts() {
	if p.ID == 0 {
		fmt.Println("no primary key")
		return
	}
	db.MyDb.Delete(p)
}
func GetAll() []Products {
	allData := make([]Products, 0)
	db.MyDb.Find(&allData)
	return allData
}

func (p *Products) FindProducts() []Products {
	findData := make([]Products, 0)
	db.MyDb.Where(p).Find(&findData)
	return findData
}
func (p *Products) UpDateProducts(m Products) {
	db.MyDb.Model(p).Save(m)
}
