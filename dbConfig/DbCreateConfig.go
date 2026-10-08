package dbConfig

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	. "pasMan/types"
)

func databasePath() string {
	_, thisFile, _, _ := runtime.Caller(0)

	projectRoot := filepath.Dir(filepath.Dir(thisFile))
	return filepath.Join(projectRoot, "database", "database.db")
}

var DB *gorm.DB

func InitDB() {

	dbPath := databasePath()
	databaseDir := filepath.Dir(dbPath)

	if err := os.MkdirAll(databaseDir, 0755); err != nil {
		log.Fatal("create database directory:", err)
	}

	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	err = DB.AutoMigrate(&Password{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	log.Println("Database is ready with GORM")
}


