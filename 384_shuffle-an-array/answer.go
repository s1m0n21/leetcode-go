/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/shuffle-an-array/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/22 11:42 上午
 */

package _shuffle_an_array

import "math/rand"

type Solution struct {
	n        int
	original []int
	nums     []int
}

func Constructor(nums []int) Solution {
	return Solution{
		n:        len(nums),
		original: append([]int(nil), nums...),
		nums:     nums,
	}
}

func (this *Solution) Reset() []int {
	copy(this.nums, this.original)
	return this.nums
}

func (this *Solution) Shuffle() []int {
	for i := 0; i < this.n; i++ {
		r := rand.Intn(this.n-i) + i
		this.nums[i], this.nums[r] = this.nums[r], this.nums[i]
	}
	return this.nums
}
