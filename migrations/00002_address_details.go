package migrations

import (
	// "context"
	// "database/sql"
	// "fmt"

	// "github.com/aayushkhosla/Mini-Social-Network/database"
	// "github.com/aayushkhosla/Mini-Social-Network/models"
	"context"
	"database/sql"
	"fmt"

	// "github.com/aayushkhosla/Mini-Social-Network/database"
	"github.com/aayushkhosla/Mini-Social-Network/models"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	goose.AddMigrationContext(upAddressDetails, downAddressDetails)
}

func upAddressDetails(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Test ", db)
	return db.Migrator().CreateTable(&models.AddressDetail{})
}

func downAddressDetails(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Test ", db)
	return db.Migrator().DropTable(&models.AddressDetail{})
}
