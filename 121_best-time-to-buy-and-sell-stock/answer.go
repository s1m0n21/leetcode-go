/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/13 8:01 PM
 */

package _best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	var ans int
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}

		if prices[i]-min > ans {
			ans = prices[i] - min
		}
	}

	return ans
}
