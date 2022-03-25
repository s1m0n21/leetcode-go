/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/25 16:56
 */

package _minimum_number_of_vertices_to_reach_all_nodes

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	type input struct {
		n     int
		edges [][]int
	}

	testCase := utils.NewTestCase(t, func(i input) []int {
		return findSmallestSetOfVertices(i.n, i.edges)
	})

	testCase.SetAndRun(input{6, [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}}, []int{0, 3})
	testCase.SetAndRun(input{5, [][]int{{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4}}}, []int{0, 2, 3})
}
