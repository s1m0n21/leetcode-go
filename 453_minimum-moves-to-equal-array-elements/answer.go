/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/20 10:26 上午
 */

package _minimum_moves_to_equal_array_elements

func minMoves(nums []int) int {
	min := nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += nums[i]-min
	}

	return ans
}
