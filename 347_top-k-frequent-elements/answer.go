/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/top-k-frequent-elements/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/25 8:10 下午
 */

package _top_k_frequent_elements

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	var freq = make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	var h = &Heap{}
	heap.Init(h)
	for n, f := range freq {
		heap.Push(h, [2]int{n, f})
	}

	var res = make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(h).([2]int)[0]
	}

	return res
}

type Heap [][2]int

func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.([2]int))
}

func (h *Heap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func (h Heap) Less(i, j int) bool {
	return h[i][1] > h[j][1]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) Len() int {
	return len(h)
}