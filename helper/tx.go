package helper

import (
	"database/sql"
	"log"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		errRollback := tx.Rollback()
		if err != nil {
			log.Fatal(errRollback)
		}
		panic(err)
	} else {
		errCommit := tx.Commit()
		if err != nil {
			log.Fatal(errCommit)
		}
	}
}
