/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/card-flipping-game/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/31 12:01 下午
 */

package _card_flipping_game

import "math"

func flipgame(fronts []int, backs []int) int {
	m := make(map[int]struct{})

	for i := 0; i < len(fronts); i++ {
		m[fronts[i]] = struct{}{}
		m[backs[i]] = struct{}{}
	}

	for i := 0; i < len(fronts); i++ {
		if fronts[i] == backs[i] {
			delete(m, fronts[i])
		}
	}

	if len(m) == 0 {
		return 0
	}

	min := math.MaxInt32
	for k := range m {
		if k < min {
			min = k
		}
	}

	return min
}
