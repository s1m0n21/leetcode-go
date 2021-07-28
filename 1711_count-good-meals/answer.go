/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/count-good-meals/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/7 5:26 下午
 */

package _count_good_meals

func countPairs(deliciousness []int) int {
	var ans = 0
	var max = deliciousness[0]
	for _, n := range deliciousness[1:] {
		if n > max {
			max = n
		}
	}

	max = max * 2
	cnt := make(map[int]int)
	for _, n := range deliciousness {
		for sum := 1; sum <= max; sum <<= 1 {
			ans += cnt[sum-n]
		}
		cnt[n]++
	}

	return ans % (1e9 + 7)
}
