package model

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	User		User	`gorm:"foreignkey:UserId"`
	Product		Product	`gorm:"foreignkey:ProductID"`
	UserID		uint
	ProductID	uint
	Quantity	int		`json:"quantity"`
}

func (Order) TableName() string {
	return "orders"
}