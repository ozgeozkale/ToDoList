package Config

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var err error

// Creates connection with database.
func InitDB() {
	DB, err = sql.Open("sqlite3", "./todo.db")
	if err != nil {
		panic("Failed while connecting to database!")
	} else {
		fmt.Printf("Connected to database.\n")
	}
}
