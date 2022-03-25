/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/minimum-number-of-vertices-to-reach-all-nodes/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/25 16:53
 */

package _minimum_number_of_vertices_to_reach_all_nodes

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	var ans []int
	inDegree := make([]int, n)
	for _, edge := range edges {
		inDegree[edge[1]]++
	}

	for i, ind := range inDegree {
		if ind == 0 {
			ans = append(ans, i)
		}
	}

	return ans
}
