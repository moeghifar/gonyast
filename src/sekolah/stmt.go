package sekolah

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/moeghifar/gonyast/src/util"
)

var (
	stmtGetSekolah *sqlx.Stmt
)

// InitStatement ...
func InitStatement() (err error) {
	// Prepared statement here
	stmtGetSekolah, err = util.DBCon.Preparex(`SELECT * FROM tbl_sklh ORDER BY id DESC`)
	if err != nil {
		log.Println("[ERROR] Failed prepared stmtGetSekolah query :: `SELECT * FROM tbl_sklh` ->", err)
	}
	return
}
