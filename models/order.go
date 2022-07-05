package models

import "time"

type Order struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	ProductRefer int     `json:"product_id"`
	Product      Product `json:"product" gorm:"foreignkey:ProductRefer"`
	UserRefer    int     `json:"user_id"`
	User         User    `json:"user" gorm:"foreignkey:UserRefer"`
}
