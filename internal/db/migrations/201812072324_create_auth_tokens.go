package migrations

import (
	"time"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var CreateAuthTokens = &gormigrate.Migration{
	ID: "201812072324",
	Migrate: func(tx *gorm.DB) error {
		type AuthToken struct {
			gorm.Model

			UserID  uint
			Token   string
			Expires time.Time
		}
		return tx.AutoMigrate(&AuthToken{}).Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.DropTable("auth_tokens").Error
	},
}
