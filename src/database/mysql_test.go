package database

import (
	"database/sql"
	"fmt"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestMySQL(t *testing.T){
	var MySQLConf  MySQLConn
	conn, err := NewConn(MySQLConf)
	if err != nil{
		t.Errorf("创建连接失败 %s", err)
	}
	SqlList := []string{ContainTable, AccountTable, PolicyTable}

	for _, sql := range(SqlList){
		fmt.Println()
		_, err := conn.Exec(sql)
		if err != nil{
			t.Errorf("SQL执行异常%s", err)
			continue
		}
	}

	rows, err := conn.Query("desc policy_stat;")
	if err != nil{
		t.Errorf("SQL执行异常%s", err)
	}

	result := queryResult(rows)
	for _, res := range(result){
		for k, v := range(res){
			fmt.Printf("k: %s, v:%s \n", k, v)	
		}
	}
}

func queryResult(rows *sql.Rows) []map[string]string{
	result := make([]map[string]string, 0, 100)
	defer func(rows *sql.Rows){
		if rows != nil{
			rows.Close()
		}
	}(rows)

	columnName, err := rows.Columns()
	if err != nil{
		fmt.Printf("row.Columns exec error %s", err)
		return result
	}

	values := make([]sql.RawBytes, len(columnName))
	scanArgs := make([]interface{}, len(values))

	for i := range values{
		scanArgs[i] = &values[i]
	}

	for rows.Next(){
		err = rows.Scan(scanArgs...)
		if err != nil{
			fmt.Println("rows.Scan err")
		}

		rowMap := make(map[string]string)
		for i, col := range values{
			if col == nil{
				rowMap[columnName[i]] = "NULL"
			}else {
				rowMap[columnName[i]] = string(col)
			}
		}
		result = append(result, rowMap)
	}
	err = rows.Err()
	if err != nil{
		fmt.Println("rows.Err have error :", err)
	}
	return result
}