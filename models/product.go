package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Product struct {
	ProductID   int       `json:"id" gorm:"primarykey, autoincrement"`
	Name        string    `json:"name" gorm:"name"`
	Price       int       `json:"price" gorm:"price"`
	Quantity    int       `json:"quantity" gorm:"quantity"`
	Description string    `json:"description" gorm:"description"`
	Image       string    `json:"image" gorm:"Image"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime:milli"`
	DeletedAt   time.Time `json:"deletedAt" gorm:"autoCreateTime"`
}
