/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/7/20 10:21
 */

package _shift_2d_grid

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	type input struct {
		grid [][]int
		k    int
	}

	tests := utils.NewTestCase(t, func(i input) [][]int {
		return shiftGrid(i.grid, i.k)
	})

	tests.SetAndRun(input{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1}, [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}})
	tests.SetAndRun(input{[][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}, 4}, [][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}})
	tests.SetAndRun(input{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 9}, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}
