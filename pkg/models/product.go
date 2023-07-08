package models

import "time"

type Product struct {
	ID                      uint      `gorm:"primaryKey"`
	UserID                  int       `gorm:"column:user_id"`
	ProductName             string    `gorm:"column:product_name"`
	ProductDescription      string    `gorm:"column:product_description"`
	ProductImages           string  `gorm:"column:product_images;type:VARCHAR"`
	CompressedProductImages string  `gorm:"column:compressed_product_images;type:VARCHAR"`
	ProductPrice            float64   `gorm:"column:product_price"`
	CreatedAt               time.Time `gorm:"column:created_at"`
	UpdatedAt               time.Time `gorm:"column:updated_at"`
}
