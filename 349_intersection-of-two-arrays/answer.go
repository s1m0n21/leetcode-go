/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/intersection-of-two-arrays/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/30 11:30 下午
 */

package _intersection_of_two_arrays

func intersection(nums1 []int, nums2 []int) []int {
	var appeared = make(map[int]int)
	var out []int

	for _, n := range nums1 {
		if _, has := appeared[n]; !has {
			appeared[n] = 1
		}
	}

	for _, n := range nums2 {
		if _, has := appeared[n]; has {
			out = append(out, n)
			delete(appeared, n)
		}
	}

	return out
}
