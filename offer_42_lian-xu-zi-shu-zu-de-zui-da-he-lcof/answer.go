/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/17 2:36 上午
 */

package offer_42_lian_xu_zi_shu_zu_de_zui_da_he_lcof

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}
