/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/13 8:01 PM
 */

package _best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	minPrice, maxProfit := 10001, 0
	for _, price := range prices {
		maxProfit = max(maxProfit, price-minPrice)
		minPrice = min(price, minPrice)
	}

	return maxProfit
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
