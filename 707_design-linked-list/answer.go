/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/design-linked-list/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/8 2:51 下午
 */

package _design_linked_list

type listNode struct {
	val        int
	prev, next *listNode
}

type MyLinkedList struct {
	length int
	head   *listNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= this.length {
		return -1
	}

	curr := this.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}

	return curr.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	node := &listNode{val: val, next: this.head}
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node
	this.length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	node := &listNode{val: val}

	curr := this.head
	if curr == nil {
		this.head = node
	} else {
		for curr.next != nil {
			curr = curr.next
		}

		curr.next = node
		node.prev = curr
	}

	this.length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	} else if index == this.length {
		this.AddAtTail(val)
		return
	} else if index > this.length {
		return
	}

	curr := this.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}

	node := &listNode{next: curr, prev: curr.prev, val: val}
	curr.prev.next = node
	curr.prev = node
	this.length++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.length {
		return
	}

	del := this.head
	for i := 0; i < index; i++ {
		del = del.next
	}

	if del.prev != nil {
		del.prev.next = del.next
		if del.next != nil {
			del.next.prev = del.prev
		}
	} else {
		this.head = this.head.next
		if this.head != nil {
			this.head.prev = nil
		}
	}
	this.length--
}
