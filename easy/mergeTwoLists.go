package main

import "fmt"

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4

type ListNode struct {
	Val  int
	Next *ListNode
}

func New() *ListNode {
	return &ListNode{0, nil}
}

// 第i位置插入元素Val
// 1->2->4  --> 1, 2->4  --> 1->5->2->4
func (head *ListNode) Insert(i int, Val int) bool {
	p := head
	j := 1
	if p != nil && j < i {
		p = p.Next
		j++
	}

	if p == nil || i < j {
		fmt.Printf("pls check the linked list or %d\n", i)
		return false
	}

	s := &ListNode{Val: Val}
	p.Next = s
	s.Next = p.Next
	return true
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

}
