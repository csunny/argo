package main

import (
	"github.com/csunny/argo/src/common"
	"fmt"
)

func main()  {

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

	common.Arr()

	fmt.Println(s)
	fmt.Println(r)
}
