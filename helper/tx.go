package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		if tx != nil {
			errorRollback := tx.Rollback()
			PanicIfError(errorRollback)
		}
		panic(err)
	} else {
		if tx != nil {
			errorCommit := tx.Commit()
			PanicIfError(errorCommit)
		}
	}
}