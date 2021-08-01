/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : anwer_test.go
 * @Date           : 2021/8/1 2:51 下午
 */

package _reverse_words_in_a_string

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct{
		input string
		expect string
	}{
		{"the sky is blue", "blue is sky the"},
		{"   hello world    ", "world hello"},
		{"a good   example", "example good a"},
		{"test", "test"},
	}

	for _, test := range tests {
		if actual := reverseWords(test.input); actual != test.expect {
			t.Errorf("input = %s, expect = %s, actual = %s", test.input, test.expect, actual)
		}
	}
}
