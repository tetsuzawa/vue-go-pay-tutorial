package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
)

// Conn - sql connection handler
var Conn *sql.DB

// NewSQLHandler - init sql handler
func init() {
	err := godotenv.Load() //Load env.file
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to load .env file at godotenv.Load()"))
	}

	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	name := os.Getenv("MYSQL_DATABASE")

	dbconf := fmt.Sprintf("%s:%s@/%s", user, pass, name)
	conn, err := sql.Open("mysql", dbconf)
	if err != nil {
		panic(err.Error)
	}
	Conn = conn
}
