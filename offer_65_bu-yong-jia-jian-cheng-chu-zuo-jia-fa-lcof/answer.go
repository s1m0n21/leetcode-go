/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/29 3:10 下午
 */

package offer_65_bu_yong_jia_jian_cheng_chu_zuo_jia_fa_lcof

func add(a int, b int) int {
	if a & b == 0 {
		return a | b
	}

	return add(a ^ b, (a & b) << 1)
}
