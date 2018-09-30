package database

import (
	"fmt"
	"testing"
)

func TestMongo(t *testing.T){
	fmt.Println("*************分割线**********")
	fmt.Println("Hello This is mongodb test file!")

	db, err := NewMongoConn()
	if err != nil{
		t.Errorf("连接Mongodb数据库异常:%s", err)
	}
	
	fmt.Println(db)
}