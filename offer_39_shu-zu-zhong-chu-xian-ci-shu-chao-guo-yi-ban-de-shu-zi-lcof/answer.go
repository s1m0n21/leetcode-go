/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/12/8 7:10 PM
 */

package offer_39_shu_zu_zhong_chu_xian_ci_shu_chao_guo_yi_ban_de_shu_zi_lcof

func majorityElement(nums []int) int {
	var vote int
	x := nums[0]

	for _, n := range nums {
		if vote == 0 {
			x = n
		}

		if n == x {
			vote++
		} else {
			vote--
		}
	}

	return x
}
