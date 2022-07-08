/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/minimum-cost-to-move-chips-to-the-same-position/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/7/8 09:27
 */

package _minimum_cost_to_move_chips_to_the_same_position

func minCostToMoveChips(position []int) int {
	var odd, event int

	for _, n := range position {
		if (n)%2 == 0 {
			event++
		} else {
			odd++
		}
	}

	if odd < event {
		return odd
	}
	return event
}
