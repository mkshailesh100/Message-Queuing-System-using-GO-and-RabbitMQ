package db

import (
	"github.com/mkshailesh100/message-queue-system/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
    dsn := "host=localhost user=user1 password=user@123 dbname=product_management port=5432 sslmode=disable"
    return gorm.Open(postgres.Open(dsn), &gorm.Config{
    })
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		return err
	}

	return nil
}
