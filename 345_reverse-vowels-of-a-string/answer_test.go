/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/29 2:23 上午
 */

package _reverse_vowels_of_a_string

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"hello", "holle"},
		{"leetcode", "leotcede"},
		{"aeiuo", "ouiea"},
	}

	for _, test := range tests {
		if actual := reverseVowels(test.input); actual != test.expect {
			t.Errorf("input = %s, expect = %s, actual = %s", test.input, test.expect, actual)
		}
	}
}
