package storage

import (
	"UrlShortner/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDbInstanse(c config.Config) *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		panic("cannot connect to database ")
	}
	return db
}
