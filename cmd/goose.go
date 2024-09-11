package main
import (
	"flag"
	"fmt"
	"log"
	"os"
	_ "github.com/aayushkhosla/Mini-Social-Network/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)
var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
)
func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()
	fmt.Println(len(args))
	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]

	dsn := "postgres://postgres:aayush@localhost/myproject?sslmode=disable"
	db, err := goose.OpenDBWithDriver("postgres", dsn)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 2 {
		arguments = append(arguments, args[2:]...)
	}

	dir := "/home/ubuntu/Desktop/Projects/Mini_Social_Network/migrations"
	if err := goose.Run(command, db, dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}