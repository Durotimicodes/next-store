package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Id              int    `json:"id" gorm:"primarykey, autoincrement"`
	Name            string `json:"name" gorm:"name"`
	Price           int    `json:"price" gorm:"price"`
	Quantity        int    `json:"quantity" gorm:"quantity"`
	ProductCategory string `json:"product_category" gorm:"product_category"`
	Description     string `json:"description" gorm:"description"`
	Image           string `json:"image" gorm:"Image"`
}

//type ProductModel struct {
//}

//func (product *ProductModel) FindAll() ([]Product, error) {
//	db, err := database.DB()
//
//	if err != nil {
//		return nil, err
//	} else {
//		var products []Product
//		var product Product
//		for rows.Next() {
//
//		}
//	}
//}
