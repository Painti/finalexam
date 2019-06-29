package database

import (
	"fmt"
)

func (conn *DBConnection) Delete(id int) (err error) {
	query := `DELETE FROM ` + conn.table + ` WHERE id=$1`
	stmt, pErr := conn.db.Prepare(query)
	if pErr != nil {
		err = pErr
		return
	}
	result, eErr := stmt.Exec(id)
	if eErr != nil {
		err = eErr
	}

	row, rErr := result.RowsAffected()
	if rErr != nil {
		err = rErr
	}
	if row == 0 {
		err = fmt.Errorf("The customer ID:%d is not exists.", id)
	}
	return
}
