/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/subarray-product-less-than-k/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/11 4:51 下午
 */

package _subarray_product_less_than_k

func numSubarrayProductLessThanK(nums []int, k int) int {
	l, r := 0, 0
	sum := 1
	ans := 0
	for ; r < len(nums); r++ {
		sum *= nums[r]
		for l <= r && sum >= k {
			sum /= nums[l]
			l++
		}
		ans += r - l + 1
	}
	return ans
}
