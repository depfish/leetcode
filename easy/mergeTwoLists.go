// 参考
// https://blog.csdn.net/s15738841819/article/details/84112069https://blog.csdn.net/s15738841819/article/details/84112069
// https://studygolang.com/articles/11690
package main

import "fmt"

func main() {
	linkedlist1 := New()
	linkedlist1.Insert(1, 4)
	linkedlist1.Insert(1, 2)
	linkedlist1.Insert(1, 1)
	linkedlist1.Traverse()

	fmt.Printf("-----------------------------\n")

	linkedlist2 := New()
	linkedlist2.Insert(1, 4)
	linkedlist2.Insert(1, 3)
	linkedlist2.Insert(1, 1)
	linkedlist2.Traverse()

	fmt.Printf("-----------------------------\n")

	linkedlist3 := mergeTwoLists(linkedlist1, linkedlist2)
	linkedlist3.Traverse()

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
func (head *ListNode) Insert(i int, val int) bool {
	p := head
	j := 1
	if p != nil && j < i {
		p = p.Next
		j++
	}

	if p == nil || j > i {
		fmt.Printf("pls check the linked list or %d\n", i)
		return false
	}

	s := &ListNode{Val: val}
	s.Next = p.Next
	p.Next = s
	return true
}

func (head *ListNode) Traverse() {
	point := head.Next
	//fmt.Println(point.Val)
	//fmt.Println(point.Next.Val)
	//fmt.Println(point.Next.Next.Val)
	for point != nil {
		fmt.Println(point.Val)
		point = point.Next
	}

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		head, tail, l1 = l1, l1, l1.Next
	} else {
		head, tail, l2 = l2, l2, l2.Next
	}

	// Till any of the linked list end
	for l1 != nil && l2 != nil {
		//找到下一个节点,添加到新链表的尾
		if l1.Val < l2.Val {
			tail.Next, l1 = l1, l1.Next
		} else {
			tail.Next, l2 = l2, l2.Next
		}
		//更新tail
		tail = tail.Next
	}

	//剩下的节点字节拼接到新链表尾部
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	//返回新链表的头指针
	return head
}
