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
	var str = ""
	var i = n
	for i != nil {
		if str == "" {
			str = fmt.Sprintf("%d", i.Val)
		} else {
			str = fmt.Sprintf("%s -> %d", str, i.Val)
		}
		i = i.Next
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
