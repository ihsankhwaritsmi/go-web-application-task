package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	//create
	db.Create(&Product{Code: "D42", Price: 100})

	//read
	var product []Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "D42")

	//update
	db.Model(&product).Update("price", 200)

	//update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	db.Model(&product).Updates(map[string]interface{}{"price": 200, "code": "F42"})
}
