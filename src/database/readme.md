### This is a repo just for database connect and operate.

Use this repo, you can connect database easily.  Maybe it's useful for beginner.

### Docker-compose 
Docker-compose.yml file can start mongodb and mysql container, and you can connect them and do some useful things.


ok, let's begin!

### MySQl Operate
The operate about mysql are show in the file `mysql.go` and `mysql_test.go` is the test file. 
In this code, I create three table which come from the `openstack swift`. 

- Run Docker Container
in this folder, i use docker-compose run mysql and mongodb container, you can run those just run 
```
docker-compose up -d
```
and then, you can run the test, just run 
```
go test
```


This is just for someone beginner or those people who is beginner for go, or intersted in golang.
Ok, Let's make go better together!