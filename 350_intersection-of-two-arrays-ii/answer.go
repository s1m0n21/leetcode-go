/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/30 11:43 下午
 */

package _intersection_of_two_arrays_ii

func intersect(nums1 []int, nums2 []int) []int {
	var appeared = make(map[int]int)
	var out []int

	for _, n := range nums1 {
		appeared[n]++
	}

	for _, n := range nums2 {
		if freq, has := appeared[n]; has && freq > 0 {
			out = append(out, n)
			appeared[n]--
		}
	}

	return out
}
