/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/10 11:01 上午
 */

package _teemo_attacking

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		timeSeries []int
		duration   int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{[]int{1, 4}, 2}, 4},
		{input{[]int{1, 2}, 2}, 3},
		{input{[]int{1, 2, 2}, 2}, 3},
	}

	for i, test := range tests {
		if actual := findPoisonedDuration(test.input.timeSeries, test.input.duration); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
