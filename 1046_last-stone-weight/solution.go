/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/last-stone-weight/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/10 09:59
 */

package _last_stone_weight

import (
	"container/heap"
)

func lastStoneWeight(stones []int) int {
	h := maxHeap{}
	heap.Init(&h)

	for _, n := range stones {
		heap.Push(&h, n)
	}

	for h.Len() > 1 {
		x := heap.Pop(&h).(int)
		y := heap.Pop(&h).(int)
		if x != y {
			heap.Push(&h, x-y)
		}
		if h.Len() == 0 {
			heap.Push(&h, 0)
		}
	}

	return heap.Pop(&h).(int)
}

type maxHeap []int

func (h *maxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }

func (h *maxHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h maxHeap) Len() int { return len(h) }
