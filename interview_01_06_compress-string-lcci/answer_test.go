/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/21 3:36 下午
 */

package interview_01_06_compress_string_lcci

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"aabcccccaaa", "a2b1c5a3"},
		{"abbccd", "abbccd"},
		{"aabbccdd", "aabbccdd"},
		{"", ""},
		{"a", "a"},
		{"rrrrrLLLLLPPPPPPRRRRRgggNNNNNVVVVVVVVVVDDDDDDDDDDIIIIIIIIIIlllllllAAAAqqqqqqqbbbNNNNffffff", "r5L5P6R5g3N5V10D10I10l7A4q7b3N4f6"},
	}

	for i, test := range tests {
		if actual := compressString(test.input); actual != test.expect {
			t.Errorf("%d: input = %s, expect = %s, actual = %s", i, test.input, test.expect, actual)
		}
	}
}
