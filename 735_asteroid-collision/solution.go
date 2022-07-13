/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/asteroid-collision/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/7/13 09:21
 */

package _asteroid_collision

func asteroidCollision(asteroids []int) []int {
	var stack []int

	for _, n := range asteroids {
		destroyed := false
		if n < 0 {
			for len(stack) > 0 {
				prev := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if prev < 0 || prev > abs(n) {
					destroyed = prev > abs(n)
					stack = append(stack, prev)
					break
				}

				if abs(n) == prev {
					destroyed = true
					break
				}
			}
		}
		if !destroyed {
			stack = append(stack, n)
		}
	}

	return stack
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
