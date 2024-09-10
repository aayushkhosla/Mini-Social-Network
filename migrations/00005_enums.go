package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateEnum, downCreateEnum)
}

func upCreateEnum(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TYPE gender AS ENUM ('male', 'female', 'other');
		CREATE TYPE marital_status AS ENUM ('single', 'married');
	`)
	if err != nil {
		return err
	}

	return nil
}

// downCreateEnum drops the ENUM types for 'gender' and 'marital_status'.
func downCreateEnum(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TYPE IF EXISTS gender;
		DROP TYPE IF EXISTS marital_status;
	`)
	if err != nil {
		return err
	}
	return nil
}





