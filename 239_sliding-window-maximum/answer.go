/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/sliding-window-maximum/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/12/3 9:16 AM
 */

package _sliding_window_maximum

func maxSlidingWindow(nums []int, k int) []int {
	var queue []int
	push := func(i int) {
		for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	ans := make([]int, 1, len(nums)-k+1)
	ans[0] = nums[queue[0]]

	for i := k; i < len(nums); i++ {
		push(i)
		for queue[0] <= i-k {
			queue = queue[1:]
		}
		ans = append(ans, nums[queue[0]])
	}

	return ans
}
