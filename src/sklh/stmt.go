package sklh

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/moeghifar/nyastmach/src/util"
)

var (
	stmtGetSklh *sqlx.Stmt
	dbc         *sqlx.DB
)

// InitStatement ...
func InitStatement() (err error) {
	dbc, err = util.InitDatabase()
	if err != nil {
		log.Println("[ERROR] Failed init database ->", err)
	}
	// Prepared statement here
	stmtGetSklh, err = dbc.Preparex(`SELECT * FROM tbl_sklh ORDER BY id DESC`)
	if err != nil {
		log.Println("[ERROR] Failed prepared stmtGetSklh query :: `SELECT * FROM tbl_sklh` ->", err)
	}
	return
}
