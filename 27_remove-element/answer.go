/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/remove-element/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/25 2:54 下午
 */

package _remove_element

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
