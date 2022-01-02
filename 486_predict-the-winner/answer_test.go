/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/12/29 9:50 AM
 */

package _predict_the_winner

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect bool
	}{
		{[]int{1, 5, 2}, false},
		{[]int{1, 5, 233, 7}, true},
	}

	for i, test := range tests {
		if actual := PredictTheWinner(test.input); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %t, actual = %t", i, test.input, test.expect, actual)
		}
	}
}
