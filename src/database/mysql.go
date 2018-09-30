package database

import (
	"fmt"
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
// MySQLConn is the config struct 
type MySQLConn struct{
	Host string
	Port int
	User string
	Dbname string
	Password string
	ChartSet string
}

// NewConn 创建一个MySql连接。
func NewConn(conn MySQLConn) (*sql.DB, error){
	conn.Host = "127.0.0.1"
	conn.Port = 3316
	conn.User = "root"
	conn.Password = "123456"
	conn.Dbname = "swift"
	conn.ChartSet = "utf8"
	db, err := connect(&conn)
	if err != nil{
		return nil, fmt.Errorf("sql connect error. %s", err)
	}
	return db, nil
}

func connect(conn *MySQLConn) (*sql.DB, error){
	db, err := sql.Open("mysql", conn.User+":"+conn.Password+"@tcp("+conn.Host+":"+strconv.Itoa(conn.Port)+")/"+conn.Dbname+"?charset="+conn.ChartSet)
	if err != nil{
		return nil, err
	}	

	if err := db.Ping(); err != nil{
		return nil, err
	}
	return db, nil
}

// CloseConn 关闭MySQL连接
func CloseConn(db *sql.DB){
	db.Close()
}

// KillQuerySession 杀死查询Session
func KillQuerySession(db *sql.DB, sessionId int64) error{
	_, err := db.Exec(fmt.Sprintf("kill query %d", sessionId))
	return err
}

// KillSession 杀死Session
func KillSession(db *sql.DB, sessionId int64) error{
	_, err := db.Exec(fmt.Sprintf("kill %d", sessionId))
	return err
}