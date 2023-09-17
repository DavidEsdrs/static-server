package connections

import (
	"github.com/DavidEsdrs/template-server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var connection *gorm.DB

func Connect() (*gorm.DB, error) {
	if connection != nil {
		return connection, nil
	}
	dsn := "file::memory:?cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	connection = db
	Migrate(connection)
	return db, nil
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Contact{})
}
