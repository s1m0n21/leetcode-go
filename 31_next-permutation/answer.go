/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/next-permutation/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/12/13 2:19 PM
 */

package _next_permutation

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}

	if i != -1 {
		j := i + 1
		for ; j < len(nums); j++ {
			if nums[j] <= nums[i] {
				break
			}
		}
		nums[i], nums[j-1] = nums[j-1], nums[i]
	}

	j := i + 1
	k := len(nums) - 1
	for j < k {
		nums[j], nums[k] = nums[k], nums[j]
		j++
		k--
	}
}
