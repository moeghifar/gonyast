package util

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // this blank import required by `jmoiron/sqlx`
)

// InitDatabase ...
func InitDatabase() (DBConn *sqlx.DB, err error) {
	DBConn, err = sqlx.Connect("postgres", Config.DBConfig["core"])
	return DBConn, err
}
