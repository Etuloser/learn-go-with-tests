package main

import (
	"fmt"

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
		panic("fail to connect database")
	}

	// 迁移 schema
	// db.AutoMigrate(&Product{})

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	// db.First(&product, 1)
	db.First(&product, "code = ?", "D42")
	db.Model(&product).Update("Price", 200)
	fmt.Println(product)
}
