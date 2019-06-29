package database

import (
	"database/sql"
	"strings"
)

func (conn *DBConnection) ReadAllData(props ...string) (rows *sql.Rows, err error) {
	fields := "*"
	if len(props) > 0 {
		fields = strings.Join(props, ",")
	}
	query := `SELECT ` + fields + ` FROM ` + conn.table
	stmt, pErr := conn.db.Prepare(query)
	if pErr != nil {
		err = pErr
		return
	}
	rows, qErr := stmt.Query()
	if qErr != nil {
		err = qErr
		return
	}
	return
}

func (conn *DBConnection) ReadByID(id int, props ...string) (row *sql.Row, err error) {
	fields := "*"
	if len(props) > 0 {
		fields = strings.Join(props, ",")
	}
	query := `SELECT ` + fields + ` FROM ` + conn.table + ` WHERE id=$1`
	stmt, pErr := conn.db.Prepare(query)
	if pErr != nil {
		err = pErr
		return
	}
	row = stmt.QueryRow(id)
	return
}
