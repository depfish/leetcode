package main

import "fmt"

func main() {
	n:=New()
	n.Insert()
	fmt.Println(n)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func New() *ListNode {
	return &ListNode{Val: 0, Next: nil}
}

func (n *ListNode) Insert() {
	//head :=n
	for i := 1; i < 6; i++ {
		(*n).Val=i

	}
}

//func reverseList(head *ListNode) *ListNode {
//
//}
