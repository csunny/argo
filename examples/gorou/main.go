/*
* goroutine 以及调度器的行为
*
 */

package main

import (
	// "strconv"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// main 是所有Go程序的入口
func main() {
	fmt.Println("Hello World!")
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	// wg 用来等待程序完成
	// 计数加2， 表示要等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")
	// 声明一个匿名函数，并创建一个goroutines
	go func() {
		// 在函数退出时候调用Done 来通知main函数工作已经完成。
		defer wg.Done()

		// 显示字母三次
		for count := 0; count < 3; count++ {
			r := rand.Intn(10)

			fmt.Println(r)
			time.Sleep(time.Second * 2)
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
			fmt.Println('\n')
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
			fmt.Println('\n')
		}
	}()

	wg.Wait()
}
