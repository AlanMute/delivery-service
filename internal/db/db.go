package db

import (
	"github.com/KrizzMU/delivery-service/internal/config"
	"github.com/KrizzMU/delivery-service/internal/core"
	"github.com/jinzhu/gorm"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("postgres", config.GetConnectionString())
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&core.Order{})
	db.AutoMigrate(&core.Delivery{})
	db.AutoMigrate(&core.Payment{})
	db.AutoMigrate(&core.Item{})

	return db
}
