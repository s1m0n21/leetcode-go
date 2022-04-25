/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode-cn.com/problems/random-pick-index/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/4/25 23:04
 */

package _98_random_pick_index

import "math/rand"

type Solution struct {
	idx map[int][]int
}

func Constructor(nums []int) Solution {
	idx := make(map[int][]int)
	for i, n := range nums {
		idx[n] = append(idx[n], i)
	}
	return Solution{idx}
}

func (s *Solution) Pick(target int) int {
	idx, has := s.idx[target]
	if !has {
		return -1
	}

	return idx[rand.Intn(len(idx))]
}
