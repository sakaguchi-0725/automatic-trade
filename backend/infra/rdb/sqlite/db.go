package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLite(filePath string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(filePath))
}
