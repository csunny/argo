package main

import (
	"github.com/csunny/argo/src/common"
	"fmt"
	"time"
)

func main() {

	start := time.Now().UnixNano() / 1e6
	res := common.Fab(45)

	//res := common.Fab_Hash(45)
	end := time.Now().UnixNano() / 1e6

	fmt.Printf("结果%d, 计算时间 %d", res, end-start)

	// 邻接矩阵
	s := [][]int{
		{0, 1, 1, 1, 0, 0},
		{1, 0, 0, 0, 1, 0},
		{1, 0, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0},
	}

	// 邻接表
	r := make(map[string][]string)
	r["A"] = []string{"B", "C", "D"}
	r["B"] = []string{"A", "E"}
	r["C"] = []string{"A", "E"}
	r["D"] = []string{"A"}
	r["E"] = []string{"B", "C", "F"}
	r["F"] = []string{"E"}

	fmt.Println(s)
	fmt.Println(r)
}