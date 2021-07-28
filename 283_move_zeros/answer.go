/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/move-zeroes/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/25 1:40 下午
 */

package _move_zeros

func moveZeroes(nums []int)  {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}

	for ; k < len(nums); k++ {
		nums[k] = 0
	}
}
