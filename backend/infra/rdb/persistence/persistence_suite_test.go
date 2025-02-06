package persistence_test

import (
	"automatic-trade/backend/infra/rdb/dto"
	"log"
	"os"
	"sync"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var testDB *gorm.DB
var once sync.Once

func TestMain(m *testing.M) {
	setupTestDB()
	cleanupTestDB()

	code := m.Run()

	os.Exit(code)
}

func setupTestDB() {
	once.Do(func() {
		var err error
		testDB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to initialize testDB: %v", err)
		}

		err = testDB.AutoMigrate(&dto.Position{})
		if err != nil {
			log.Fatalf("failed to migrate testDB: %v", err)
		}

		log.Println("successfully initialize testDB")
	})
}

func cleanupTestDB() {
	if testDB == nil {
		log.Fatalf("testDB is not initialized")
	}

	err := testDB.Exec("DELETE FROM positions").Error
	if err != nil {
		log.Fatalf("failed to cleanup positions table: %v", err)
	}
}
