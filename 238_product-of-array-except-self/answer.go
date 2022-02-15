/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/product-of-array-except-self/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/2/15 3:45 PM
 */

package _product_of_array_except_self

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans, left, right := make([]int, n), make([]int, n), make([]int, n)

	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		ans[i] = left[i] * right[i]
	}

	return ans
}
