package db

import (
	"github.com/irth/abradolf-backend/internal/db/migrations"
	"github.com/jinzhu/gorm"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Open(driver, path string) *gorm.DB {
	db, err := gorm.Open(driver, path)
	check(err)

	check(migrations.Migrate(db))

	return db
}
