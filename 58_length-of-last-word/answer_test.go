/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/15 9:07 PM
 */

package _length_of_last_word

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}

	for i, test := range tests {
		if actual := lengthOfLastWord(test.input); actual != test.expect {
			t.Errorf("%d: input = %s, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
