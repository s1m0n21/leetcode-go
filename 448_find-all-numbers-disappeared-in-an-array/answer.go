/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/25 17:03
 */

package _find_all_numbers_disappeared_in_an_array

func findDisappearedNumbers(nums []int) []int {
	var ans []int
	n := len(nums)

	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}

	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}

	return ans
}
