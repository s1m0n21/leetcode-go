/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/14 2:50 下午
 */

package _find_the_town_judge

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		n     int
		trust [][]int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{2, [][]int{{1, 2}}}, 2},
		{input{3, [][]int{{1, 3}, {2, 3}}}, 3},
		{input{3, [][]int{{1, 3}, {2, 3}, {3, 1}}}, -1},
		{input{3, [][]int{{1, 2}, {2, 3}}}, -1},
		{input{4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}}, 3},
		{input{1, nil}, 1},
	}

	for i, test := range tests {
		if actual := findJudge(test.input.n, test.input.trust); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
