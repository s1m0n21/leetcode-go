/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/2/23 4:25 PM
 */

package _find_the_winner_of_the_circular_game

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		n, k int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{5, 2}, 3},
		{input{6, 5}, 1},
	}

	for i, test := range tests {
		if actual := findTheWinner(test.input.n, test.input.k); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
