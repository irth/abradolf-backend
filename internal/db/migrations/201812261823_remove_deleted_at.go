package migrations

import (
	"github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

var RemoveDeletedAt = &gormigrate.Migration{
	ID: "201812261823",
	Migrate: func(tx *gorm.DB) error {
		err := tx.Table("users").DropColumn("deleted_at").Error
		if err != nil {
			return err
		}

		return tx.Table("auth_tokens").DropColumn("deleted_at").Error
	},
	Rollback: func(tx *gorm.DB) error {
		type User struct {
			gorm.Model
		}
		type AuthToken struct {
			gorm.Model
		}
		return tx.AutoMigrate(&User{}, &AuthToken{}).Error
	},
}
