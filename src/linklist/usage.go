package linklist

import (
	"strconv"
	"fmt"
)

// ### 链表使用

// 一、  反转链表中间部分
// 已知链表的头，将链表从位置m、n进行逆序
func (head *LinkNode) MediumReverse(m, n int) *LinkNode {
	// 首先需要算出需要逆序的个数
	// 边界条件，m与n相等的时候，反转个数为1个相当于不反转，直接返回
	if m == n {
		return head
	}

	if m > n || m > head.GetLength() {
		panic("m, n value error!")
	}

	count := n - m + 1 // 需要反转的节点个数

	var preNode *LinkNode // 起始反转的节点

	result := head

	for i := 0; i < m; i++ {
		preNode = head
		head = head.Next
	}

	modifyTail := head
	var newHead *LinkNode

	for j := 0; j <= count; j++ {
		next := head.Next
		head.Next = newHead
		head = next
	}

	modifyTail.Next = head

	if preNode != nil {
		preNode.Next = newHead
	} else {
		result = newHead
	}

	return result
}

// 二、计算两个链表的焦点

// 三、 合并两个排好序的链表

// 四、 约瑟夫环
func (tail *LinkNode) NewJosphuseRing(num int) *LinkNode {
	// 构造一个环
	for i := 0; i < num; i++ {
		fmt.Printf("%d\n", i)

		if tail.Payload == "" {
			tail := LinkNode{Payload: strconv.Itoa(i)}
			tail.Next = &tail
			fmt.Printf("---%s\n", tail.Payload)
		} else {
			newNode := LinkNode{Payload: strconv.Itoa(i)}
			fmt.Printf("===%s\n", newNode.Payload)

			newNode.Next = tail.Next
			tail.Next = &newNode

			tail = &newNode
		}
	}

	return tail
}
