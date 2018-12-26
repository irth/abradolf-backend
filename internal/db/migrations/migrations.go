package migrations

import (
	"github.com/irth/abradolf-backend/internal/db/models"
	"github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		CreateUsers,
		CreateAuthTokens,
		RemoveDeletedAt, // DROP COLUMN is not supported by sqlite3
	})

	m.InitSchema(func(tx *gorm.DB) error {
		return tx.AutoMigrate(&models.User{}, &models.AuthToken{}).Error
	})

	return m.Migrate()
}
