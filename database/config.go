package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbURL = "postgres://postgres:aayush@localhost/myproject?sslmode=disable"
var dbi = os.Getenv("Database")
var GORM_DB *gorm.DB
var DB_MIGRATOR gorm.Migrator


func ConnectToDatabase() (error) {
    fmt.Println("helloooooooooooooooo ---->>>>",dbi)
    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
    if err == nil {
        GORM_DB = db
        DB_MIGRATOR = db.Migrator()
    }
    return err
}


