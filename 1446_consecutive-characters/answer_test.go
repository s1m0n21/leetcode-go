/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/12/6 5:34 PM
 */

package _consecutive_characters

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{"leetcode", 2},
		{"abbcccddddeeeeedcba", 5},
		{"triplepillooooow", 5},
		{"hooraaaaaaaaaaay", 11},
		{"tourist", 1},
	}

	for i, test := range tests {
		if actual := maxPower(test.input); actual != test.expect {
			t.Errorf("%d: input = %s, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
