/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/4sum-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/6 1:39 下午
 */

package _four_sum_ii

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var freq = make(map[int]int)
	for i := 0; i < len(nums3); i++ {
		for j := 0 ; j < len(nums4); j++ {
			freq[nums3[i] + nums4[j]]++
		}
	}

	var out = 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if count, has := freq[0 - nums1[i] - nums2[j]]; has {
				out += count
			}
		}
	}

	return out
}
