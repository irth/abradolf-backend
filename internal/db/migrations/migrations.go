package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func Migrate(db *gorm.DB) error {
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		CreateUsers,
		CreateAuthTokens,
	}).Migrate()
}
