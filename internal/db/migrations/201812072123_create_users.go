package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var CreateUsers = &gormigrate.Migration{
	ID: "201812072123",
	Migrate: func(tx *gorm.DB) error {
		type User struct {
			gorm.Model

			Username     string `gorm:"unique"`
			PasswordHash []byte
		}
		return tx.AutoMigrate(&User{}).Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.DropTable("users").Error
	},
}
