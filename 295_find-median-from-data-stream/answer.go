/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/find-median-from-data-stream/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/27 9:56 上午
 */

package _find_median_from_data_stream

import (
	"container/heap"
)

type MedianFinder struct {
	min *minHeap
	max *maxHeap
}

func Constructor() MedianFinder {
	mf := MedianFinder{&minHeap{}, &maxHeap{}}
	heap.Init(mf.min)
	heap.Init(mf.max)
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.max, num)
	heap.Push(this.min, heap.Pop(this.max))

	for this.max.Len() < this.min.Len() {
		heap.Push(this.max, heap.Pop(this.min))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.max.Len() > this.min.Len() {
		return float64((*this.max)[0])
	}
	return float64((*this.max)[0] + (*this.min)[0]) / 2
}

type minHeap []int

func (h *minHeap) Push(v interface{}) { *h = append(*h, v.(int))}

func (h *minHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h minHeap) Len() int { return len(h) }

type maxHeap []int

func (h *maxHeap) Push(v interface{}) { *h = append(*h, v.(int))}

func (h *maxHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h maxHeap) Len() int { return len(h) }
