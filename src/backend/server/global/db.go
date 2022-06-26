package global

import (
	"server/internal/model"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func SetupDB() error {
	var err error
	DB, err = model.NewDBEngine(DatabaseSettings)
	if err != nil {
		return err
	}

	return nil
}
