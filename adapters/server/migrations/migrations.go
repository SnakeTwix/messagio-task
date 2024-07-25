package migrations

import (
	"gorm.io/gorm"
	"messagio/core/entity"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.Message{},
		&entity.MessageRead{},
	)

	return err
}
