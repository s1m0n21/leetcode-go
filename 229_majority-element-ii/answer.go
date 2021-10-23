/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/majority-element-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/22 5:04 下午
 */

package _majority_element_ii

func majorityElement(nums []int) []int {
	e := make(map[int]int)
	for _, n := range nums {
		e[n]++
	}

	var ans []int
	m := len(nums) / 3
	for _, n := range nums {
		if e[n] > m {
			ans = append(ans, n)
			delete(e, n)
		}
	}

	return ans
}
