package util

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // this blank import required by `jmoiron/sqlx`
)

// InitDatabase ...
func InitDatabase() (DBConn *sqlx.DB, err error) {
	DBConn, err = sqlx.Connect("postgres", "user=postgres dbname=db_alpha_app password=root sslmode=disable")
	return DBConn, err
}
