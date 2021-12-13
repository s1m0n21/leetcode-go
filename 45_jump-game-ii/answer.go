/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/jump-game-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/12/13 1:23 PM
 */

package _jump_game_ii

func jump(nums []int) int {
	n := len(nums)
	var steps = make([]int, n)
	var r int

	for i := 0; i < n; i++ {
		var nr int
		if i+nums[i] < n-1 {
			nr = i + nums[i]
		} else {
			nr = n - 1
		}

		if nr < r {
			continue
		} else {
			for j := r + 1; j <= nr; j++ {
				steps[j] = steps[i] + 1
			}
			r = nr
		}
	}

	return steps[n-1]
}
