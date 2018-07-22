package main

import (
	"github.com/csunny/argo/src/common"
	"fmt"
	"time"
)

func main() {

	start := time.Now().UnixNano()/1e6
	res := common.Fab(45)

	//res := common.Fab_Hash(45)
	end := time.Now().UnixNano()/1e6

	fmt.Printf("结果%d, 计算时间 %d", res, end-start)
}
