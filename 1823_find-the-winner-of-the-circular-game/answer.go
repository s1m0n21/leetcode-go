/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/find-the-winner-of-the-circular-game/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/2/23 4:01 PM
 */

package _find_the_winner_of_the_circular_game

type listNode struct {
	prev, next *listNode
	idx        int
}

func findTheWinner(n int, k int) int {
	var head = &listNode{idx: 1}
	tail := head
	for i := 1; i < n; i++ {
		node := &listNode{
			prev: tail,
			idx:  i + 1,
		}
		tail.next = node
		tail = node
	}
	tail.next = head
	head.prev = tail

	for head.next != head {
		for i := 1; i < k; i++ {
			head = head.next
		}
		head.prev.next = head.next
		head.next.prev = head.prev
		head = head.next
	}

	return head.idx
}
