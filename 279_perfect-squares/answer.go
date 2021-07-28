/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/perfect-squares/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/24 10:48 上午
 */

package _perfect_squares

type elem struct {
	num int
	step int
}

func numSquares(n int) int {
	var queue = make([]elem, 0)
	queue = append(queue, elem{
		num:  n,
		step: 0,
	})

	var visited = make([]bool, n + 1)
	visited[n] = true

	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]

		for i := 1; ; i++ {
			a := e.num - i * i
			if a < 0 {
				break
			}

			if a == 0 {
				return e.step + 1
			}

			if !visited[a] {
				queue = append(queue, elem{
					num:  a,
					step: e.step + 1,
				})
				visited[a] = true
			}
		}
	}

	return 0
}
