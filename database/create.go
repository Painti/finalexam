package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

func (conn *DBConnection) Create(data map[string]interface{}, props ...string) (row *sql.Row, err error) {
	i := 0
	var datas []interface{}
	var cols []string
	var value []string
	for k, v := range data {
		i++
		num := strconv.Itoa(i)
		value = append(value, "$"+num)
		cols = append(cols, k)
		datas = append(datas, v)
	}
	if i == 0 {
		return nil, fmt.Errorf("not found new data")
	}
	props = append(props, "id")
	propsFields := strings.Join(props, ",")
	colFields := strings.Join(cols, ",")
	valFields := strings.Join(value, ",")
	query := `INSERT INTO ` + conn.table + `(` + colFields + `) VALUES(` + valFields + `) RETURNING ` + propsFields
	row = conn.db.QueryRow(query, datas...)
	return
}
