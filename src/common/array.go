package common

import (
	"fmt"
)

func Arr()  {
	a := [2] int{1, 2}   // 数组， 一开始就分配好了内存，大小限定，不能拓展
	b := []int{1, 2, 3, 4, 5, 6, 7}  // 切片，大小没有限定，可以动态拓展

	for _, i := range a{
		fmt.Printf("a 数组遍历 %d\n", i)
	}

	for _, j := range b{
		fmt.Printf("b 切片遍历 %d\n", j)
	}

	//输出数组的第二个元素
	fmt.Printf("a的第二个元素 %d\n", a[1])

	// 输出切片的第五个元素
	fmt.Printf("b的第五个元素 %d\n", b[4])
}

func Insert(p int, value int, s []int) []int {
	// 给切片插入元素 p为插入位置 value 是插入值  s是原来的数组

	result := append(s[:p], value)
	result = append(result, s[p:]...)
	return result
}

func Reverse(s []int) {
	// 反转一个切片
	// 此处利用go语言多重赋值的特性
	for i, j:=0, len(s)-1; i<j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}