package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"vault/internal/database/migrations"
	"vault/internal/database/models"
)

func New(databaseURL string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		TranslateError: true,
	})
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.User{}, &models.Session{}, &models.Repository{}, &models.Event{}); err != nil {
		return err
	}

	if err := migrations.RemoveUsersUpdatedAt(db); err != nil {
		return err
	}

	return migrations.EnsureEventsConstraints(db)
}
