/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/16 9:36 上午
 */

package offer_53_zai_pai_xu_shu_zu_zhong_cha_zhao_shu_zi_lcof

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] < target {
			l++
		} else if nums[r] > target {
			r--
		}

		if l >= len(nums) || r < 0 || nums[l] == target && nums[r] == target {
			break
		}
	}

	return r - l + 1
}
