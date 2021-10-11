/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/11 4:18 下午
 */

package interview_08_05_recursive_mulitply_lcci

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		a, b int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{1, 10}, 1 * 10},
		{input{1000, 1000}, 1000 * 1000},
		{input{4, 3}, 4 * 3},
		{input{400000, 398}, 400000 * 398},
	}

	for i, test := range tests {
		if actual := multiply(test.input.a, test.input.b); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
