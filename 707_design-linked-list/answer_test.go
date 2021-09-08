/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/9/8 3:32 下午
 */

package _design_linked_list

import (
	"fmt"
	"testing"
)

func TestAnswer(t *testing.T) {
	linkedList := Constructor()

	linkedList.AddAtHead(7)
	printLinkedList(t, linkedList.head)
	linkedList.AddAtHead(2)
	printLinkedList(t, linkedList.head)
	linkedList.AddAtHead(1)
	printLinkedList(t, linkedList.head)
	linkedList.AddAtIndex(3, 0)
	printLinkedList(t, linkedList.head)
	linkedList.DeleteAtIndex(2)
	printLinkedList(t, linkedList.head)
	linkedList.AddAtHead(6)
	printLinkedList(t, linkedList.head)
	linkedList.AddAtTail(4)
	printLinkedList(t, linkedList.head)
	t.Log(linkedList.Get(4))
	linkedList.AddAtHead(4)
	printLinkedList(t, linkedList.head)
	linkedList.AddAtIndex(5, 0)
	printLinkedList(t, linkedList.head)
	linkedList.AddAtHead(6)
	printLinkedList(t, linkedList.head)
}

func printLinkedList(t *testing.T, n *listNode) {
	if n == nil {
		t.Log("NIL")
		return
	}

	var str string
	curr := n
	head := n
	loop := false

	for curr != nil {
		if len(str) == 0 {
			str = fmt.Sprintf("%d", curr.val)
		} else {
			str = fmt.Sprintf("%s -> %d", str, curr.val)
		}
		if curr.next == head {
			loop = true
			break
		}
		curr = curr.next
	}

	if loop {
		str = fmt.Sprintf("%s -> LOOP", str)
	} else {
		str = fmt.Sprintf("%s -> NIL", str)
	}

	t.Log(str)
}
