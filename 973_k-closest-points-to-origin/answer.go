/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/k-closest-points-to-origin/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/4 8:29 PM
 */

package _k_closest_points_to_origin

import "container/heap"

func kClosest(points [][]int, k int) [][]int {
	var ans [][]int

	h := make(hp, k)
	for i, p := range points[:k] {
		h[i] = pair{p[0]*p[0] + p[1]*p[1], p}
	}
	heap.Init(&h)

	for _, p := range points[k:] {
		if dist := p[0]*p[0] + p[1]*p[1]; dist < h[0].dist {
			h[0] = pair{dist, p}
			heap.Fix(&h, 0)
		}
	}

	for _, p := range h {
		ans = append(ans, p.point)
	}

	return ans
}

type pair struct {
	dist  int
	point []int
}

type hp []pair

func (h hp) Len() int { return len(h) }

func (h hp) Less(i, j int) bool { return h[i].dist > h[j].dist }

func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }

func (h *hp) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}
