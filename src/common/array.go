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
