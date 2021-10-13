/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/13 4:14 下午
 */

package _longest_substring_of_all_vowels_in_order

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{"aeiaaioaaaaeiiiiouuuooaauuaeiu", 13},
		{"aeeeiiiioooauuuaeiou", 5},
		{"a", 0},
		{"aeiou", 5},
		{"uoaeiiiou", 7},
	}

	for i, test := range tests {
		if actual := longestBeautifulSubstring(test.input); actual != test.expect {
			t.Errorf("%d: input = %s, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
