/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/12/6 4:04 PM
 */

package _x_of_a_kind_in_a_deck_of_cards

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect bool
	}{
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, true},
		{[]int{1, 1, 1, 2, 2, 2, 3, 3}, false},
		{[]int{1}, false},
		{[]int{1, 1}, true},
	}

	for i, test := range tests {
		if actual := hasGroupsSizeX(test.input); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %t, actual = %t", i, test.input, test.expect, actual)
		}
	}
}
