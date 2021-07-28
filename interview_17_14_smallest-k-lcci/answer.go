/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/smallest-k-lcci/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/26 11:33 上午
 */

package interview_17_14_smallest_k_lcci

import "sort"

func smallestK(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
