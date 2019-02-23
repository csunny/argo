package main

import (
	"fmt"
	"net"
	"time"
	// "net/url"
)

func main() {
	fmt.Println("Hello World!")

	conn, err := net.DialTimeout("tcp", "10.103.113.127:8080", 1*time.Second)

	if err != nil {
		fmt.Println(err)
		fmt.Println("err has accur")
	}

	defer conn.Close()

}
