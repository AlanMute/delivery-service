package db

import (
	"github.com/KrizzMU/delivery-service/internal/config"
	"github.com/jinzhu/gorm"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("postgres", config.GetConnectionString())
	if err != nil {
		panic(err)
	}
	return db
}
