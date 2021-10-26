/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/next-greater-element-i/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/26 9:37 上午
 */

package _next_greater_element_i

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}

	m := make(map[int]int)
	for i, v := range nums1 {
		m[v] = i
	}

	stack := []int{0}
	for i := 1; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			if idx, has := m[nums2[top]]; has {
				ans[idx] = nums2[i]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return ans
}
