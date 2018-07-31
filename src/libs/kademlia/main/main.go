package main

import "fmt"

func main()  {
	s := func(x int) int {
		return x * 2
	}

	fmt.Println(s(4))
}
