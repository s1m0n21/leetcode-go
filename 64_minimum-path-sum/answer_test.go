/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/7 4:22 PM
 */

package _minimum_path_sum

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  [][]int
		expect int
	}{
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
	}

	for i, test := range tests {
		if actual := minPathSum(test.input); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.expect, test.expect, actual)
		}
	}
}
