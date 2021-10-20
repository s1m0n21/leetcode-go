/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/20 11:47 上午
 */

package _check_if_numbers_are_ascending_in_a_sentence

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct{
		input string
		expect bool
	}{
		{"1 box has 3 blue 4 red 6 green and 12 yellow marbles", true},
		{"hello world 5 x 5", false},
		{"sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s", false},
		{"4 5 11 26", true},
	}

	for i, test := range tests {
		if actual := areNumbersAscending(test.input); actual != test.expect {
			t.Errorf("%d: input = %s, exepct = %t, actual = %t", i, test.input, test.expect, actual)
		}
	}
}
