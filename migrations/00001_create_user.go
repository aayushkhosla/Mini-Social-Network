package migrations

import (
	"context"
	"database/sql"
	"fmt"

	// "fmt"

	// "github.com/aayushkhosla/Mini-Social-Network/database"
	"github.com/aayushkhosla/Mini-Social-Network/models"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	goose.AddMigrationContext(upCreateUser, downCreateUser)
	//goose.AddMigrationContext(upAddressDetails,downAddressDetails)
}

func upCreateUser(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Test ", db)
	return db.Migrator().CreateTable(&models.User{})


	// if err:=database.DB_MIGRATOR.CreateTable(&models.User{});err !=nil {
	// 	return err
	// }

	
	// if err :=  database.DB_MIGRATOR.CreateTable(&models.AddressDetail{}) ;err !=nil {
	// 	return err
	// }
	
	// if err :=  database.DB_MIGRATOR.CreateTable(&models.Follow{}) ;err !=nil {
	// 	return err
	// }
	// if err :=  database.DB_MIGRATOR.CreateTable(&models.OfficeDetail{}) ;err !=nil {
	// 	return err
	// }
	// return nil
	//  database.DB_MIGRATOR.CreateTable(&models.Follow{})
	//  database.DB_MIGRATOR.CreateTable(&models.OfficeDetail{})
}

func downCreateUser(ctx context.Context, tx *sql.Tx) error {

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Test ", db)
	return db.Migrator().DropTable(&models.User{})
	// This code is executed when the migration is rolled back.
	// return database.DB_MIGRATOR.DropTable(&models.User{})
}


// func upAddressDetails(ctx context.Context, tx *sql.Tx) error {
// 	fmt.Println("Test ", database.DB_MIGRATOR)
// 	return database.DB_MIGRATOR.CreateTable(&models.AddressDetail{})
// }

// func downAddressDetails(ctx context.Context, tx *sql.Tx) error {
// 	// This code is executed when the migration is rolled back.
// 	return database.DB_MIGRATOR.DropTable(&models.AddressDetail{})
// }
