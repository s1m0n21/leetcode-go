/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/7 5:54 下午
 */

package offer_45_ba_shu_zu_pai_cheng_zui_xiao_de_shu_lcof

import (
	"strconv"
)

func minNumber(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}

	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if (nums[j-1] * mul(nums[j]) + nums[j]) > (nums[j] * mul(nums[j-1]) + nums[j-1]) {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}

	var res string
	for _, n := range nums {
		res += strconv.Itoa(n)
	}

	return res
}

func mul(n int) int {
	if n == 0 {
		return 10
	}

	m := 1
	for n > 0 {
		m *= 10
		n /= 10
	}
	return m
}
