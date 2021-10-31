package database

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	// ensure folder ./data exists and if not create it
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		if err := os.Mkdir("./data", 0755); err != nil {
			panic(fmt.Sprintf("[database][init] error: %s", err))
		}
	}
}

// Migrate migrates the database
func Migrate() error {
	tables := []interface{}{
		&Project{},
	}
	return db.AutoMigrate(tables...)
}

// Connect connects the database, panics if the connection failes
func Connect() {
	var err error
	db, err = gorm.Open(sqlite.Open("data/database.db"))
	if err != nil {
		panic(fmt.Sprintf("[database][init] error: %s", err))
	}
	// enable foreign key constraints
	db.Exec("PRAGMA foreign_keys = ON;")
}
