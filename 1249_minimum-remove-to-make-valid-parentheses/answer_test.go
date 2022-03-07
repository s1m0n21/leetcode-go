/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/7 11:56 AM
 */

package _minimum_remove_to_make_valid_parentheses

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"lee(t(c)o)de)", "lee(t(c)o)de"},
		{"a)b(c)d", "ab(c)d"},
		{"))((", ""},
	}

	for i, test := range tests {
		if actual := minRemoveToMakeValid(test.input); actual != test.expect {
			t.Errorf("%d: input = %s, expect = %s, actual = %s", i, test.input, test.expect, actual)
		}
	}
}
