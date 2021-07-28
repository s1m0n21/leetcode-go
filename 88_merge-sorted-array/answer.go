/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/merge-sorted-array/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/26 10:12 下午
 */

package _merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int)  {
	k, j := m-1, n-1

	for i := m+n-1; k >= 0 || j >= 0 && i >= 0; i-- {
		n := 0
		if k >= 0 && j >= 0 {
			if nums1[k] > nums2[j] {
				n = nums1[k]
				k--
			} else {
				n = nums2[j]
				j--
			}
		} else if k < 0 {
			n = nums2[j]
			j--
		} else {
			n = nums1[k]
			k--
		}

		nums1[i] = n
	}
}

