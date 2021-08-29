/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/29 3:25 下午
 */

package _count_of_matches_in_tournament

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  int
		expect int
	}{
		{7, 6},
		{14, 13},
	}

	for _, test := range tests {
		if actual := numberOfMatches(test.input); actual != test.expect {
			t.Errorf("input = %d, exepct = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
