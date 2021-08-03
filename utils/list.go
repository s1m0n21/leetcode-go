/**
 * @Author         : s1m0n21
 * @Description    : List util
 * @Project        : leetcode-go
 * @File           : list.go
 * @Date           : 2021/6/6 11:20 上午
 */

package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	if n == nil {
		return "NIL"
	}

	var str string
	curr := n
	head := n
	loop := false

	for curr != nil {
		if len(str) == 0 {
			str = fmt.Sprintf("%d", curr.Val)
		} else {
			str = fmt.Sprintf("%s -> %d", str, curr.Val)
		}
		if curr.Next == head {
			loop = true
			break
		}
		curr = curr.Next
	}

	if loop {
		str = fmt.Sprintf("%s -> LOOP", str)
	} else {
		str = fmt.Sprintf("%s -> NIL", str)
	}

	return str
}

func MakeListNode(nums ...int) *ListNode {
	var head *ListNode
	var tail *ListNode

	for _, n := range nums {
		if head == nil {
			head = &ListNode{Val: n}
			tail = head
		} else {
			tail.Next = &ListNode{Val: n}
			tail = tail.Next
		}
	}

	return head
}

func MakeCycleListNode(nums ...int) *ListNode {
	var head *ListNode
	var tail *ListNode

	for _, n := range nums {
		if head == nil {
			head = &ListNode{Val: n}
			tail = head
		} else {
			tail.Next = &ListNode{Val: n}
			tail = tail.Next
		}
	}

	tail.Next = head

	return head
}

func SameList(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}

	if a != nil || b != nil {
		return false
	}

	return true
}
