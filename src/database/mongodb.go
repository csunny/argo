// This is the mongo operate use golang. 
// In this file i will show the insert、 delete、 update、search in mongo with go

package database

import (
	"fmt"
	"context"
	"strconv"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// MongoConn is the config of mongo connect
type MongoConn struct{
	Host string
	Port int
	DbName string
}

// NewMongoConn create a new connect 
func NewMongoConn() (db *mongo.Database, err error){
	var mongoConf MongoConn
	mongoConf.Host = "127.0.0.1"
	mongoConf.Port =  27017 
	mongoConf.DbName = "swift"

	client, err := mongo.Connect(context.Background(), "mongodb" + ":" + "//" + mongoConf.Host + ":"+ strconv.Itoa(mongoConf.Port), nil)
	if err != nil{
		fmt.Printf("连接mongodb Error %s:",  err)
		return nil, err
	}
	db = client.Database("swift")
	return db, nil
}
