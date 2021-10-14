/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/find-the-town-judge/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/14 2:45 下午
 */

package _find_the_town_judge

func findJudge(n int, trust [][]int) int {
	p := make([]int, n)
	for _, t := range trust {
		p[t[0]-1] = -1
		if p[t[1]-1] == -1 {
			continue
		}
		p[t[1]-1]++
	}

	for i, t := range p {
		if t == n-1 {
			return i+1
		}
	}

	return -1
}
