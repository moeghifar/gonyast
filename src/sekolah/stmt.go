package sekolah

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/moeghifar/gonyast/src/util"
)

var (
	stmtGetSekolah *sqlx.Stmt
	dbc            *sqlx.DB
)

// InitStatement ...
func InitStatement() (err error) {
	dbc, err = util.InitDatabase()
	if err != nil {
		log.Println("[ERROR] Failed init database ->", err)
	}
	// Prepared statement here
	stmtGetSekolah, err = dbc.Preparex(`SELECT * FROM tbl_sklh ORDER BY id DESC`)
	if err != nil {
		log.Println("[ERROR] Failed prepared stmtGetSekolah query :: `SELECT * FROM tbl_sklh` ->", err)
	}
	return
}
