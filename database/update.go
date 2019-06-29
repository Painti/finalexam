package database

import (
	"fmt"
	"strconv"
	"strings"
)

func (conn *DBConnection) Update(id int, data map[string]interface{}) (err error) {
	i := 0
	datas := []interface{}{id}
	var cols []string
	for k, v := range data {
		i++
		num := strconv.Itoa(i + 1)
		cols = append(cols, k+"=$"+num)
		datas = append(datas, v)
	}
	if i == 0 {
		return fmt.Errorf("not found new data")
	}
	colFields := strings.Join(cols, ",")
	query := `UPDATE ` + conn.table + ` SET ` + colFields + ` WHERE id=$1`
	stmt, pErr := conn.db.Prepare(query)
	if pErr != nil {
		err = pErr
		return
	}
	_, eErr := stmt.Exec(datas...)
	if eErr != nil {
		err = eErr
	}

	return
}
