/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/7 12:47 PM
 */

package _coin_change

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		coins  []int
		amount int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{[]int{1, 2, 5}, 11}, 3},
		{input{[]int{2}, 3}, -1},
		{input{[]int{1}, 0}, 0},
		{input{[]int{1}, 1}, 1},
		{input{[]int{1}, 2}, 2},
	}

	for i, test := range tests {
		if actual := coinChange(test.input.coins, test.input.amount); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
