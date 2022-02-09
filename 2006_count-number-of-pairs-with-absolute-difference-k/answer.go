/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/count-number-of-pairs-with-absolute-difference-k/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/2/9 9:28 AM
 */

package _count_number_of_pairs_with_absolute_difference_k

func countKDifference(nums []int, k int) int {
	var ans int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) == k {
				ans++
			}
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
