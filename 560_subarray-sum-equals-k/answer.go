/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/subarray-sum-equals-k/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/2/14 2:07 PM
 */

package _subarray_sum_equals_k

func subarraySum(nums []int, k int) int {
	var count, pre int
	m := map[int]int{0: 1}
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if n, ok := m[pre-k]; ok {
			count += n
		}
		m[pre]++
	}
	return count
}
