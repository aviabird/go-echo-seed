package db

import (
	"fmt"

	model "github.com/aviabird/go-echo-seed/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var devDSN string = "host=localhost user=postgres password=postgres dbname=echoweb_dev port=5432 sslmode=disable"
var testDSN string = "host=localhost user=postgres password=postgres dbname=echoweb_test port=5432 sslmode=disable"
var testSavePoint = "testStart"

func New() *gorm.DB {
	db, err := gorm.Open(postgres.Open(devDSN), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect database with error: ", err)
	}

	db.Use(dbresolver.Register(dbresolver.Config{}).SetMaxIdleConns(3))
	return db
}

func TestDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(testDSN), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect database with error: ", err)
	}

	// Use to rollback to after test run
	db.SavePoint(testSavePoint)
	db.Use(dbresolver.Register(dbresolver.Config{}).SetMaxIdleConns(3))
	return db
}

func DropTestDB(db *gorm.DB) *gorm.DB {
	return db.RollbackTo(testSavePoint)
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
	)
}
