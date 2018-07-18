package linklist

import (
	"fmt"
	"log"
)

// Item 可以理解为范性，也就是任意的数据类型
type Item interface {

}

// 一个节点，除了自身的数据之外，还必须指向下一个节点，尾部节点指向为nil
type LinkNode struct {
	Payload Item    //  Payload 为任意数据类型
	Next *LinkNode
}

// go语言方法，对比与下面的NewLinkNode，方法可以理解为面向对象里面对象的方法，虽然实现的功能
// 跟函数类似，但是方式是绑定在对象上的，也就是说我们此处的Add是绑定与head这个LinkNode对象的。
// 这个是go语言区别与其他语言的设计方式，也是go语言很重要的一部分。
func (head *LinkNode) Add(payload Item)  {
	// 这里采用尾部插入的方式，给链表添加元素
	point := head

	for point.Next != nil{
		point = point.Next
	}
	newNode := LinkNode{payload, nil}
	point.Next = &newNode


	// 头部插入
	//newNode := LinkNode{payload, nil}
	//newNode.Next = head
}

// 创建一个新的链表。 函数，对比与上面的方法，函数是没有绑定任何对象的。
// go语言的函数需要指明参数跟返回值，在此函数中，我们的参数是length，返回值是一个LinkNode对象
// 除了绑定之外，函数跟方法并没有什么不同
func NewLinkNode(length int) *LinkNode  {
	if length <= 0{
		fmt.Printf("链表长度必须大于0")
		log.Panic("链表长度必须大于0")
	}
	var head *LinkNode

	head = &LinkNode{}

	for i := 0; i<length; i++{
		var newNode *LinkNode
		newNode = &LinkNode{Payload: i}
		newNode.Next = head
		head = newNode
	}
	return head
}