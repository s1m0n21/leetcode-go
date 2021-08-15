/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/15 11:56 下午
 */

package _valid_parenthesis_string

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{"()", true},
		{"*)", true},
		{"((", false},
		{"(*))", true},
		{"", true},
		{"((**", true},
		{"(*)", true},
		{"(()", false},
		{"(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())", false},
	}

	for _, test := range tests {
		if actual := checkValidString(test.input); actual != test.expect {
			t.Errorf("input = %s, expect = %t, actual = %t", test.input, test.expect, actual)
		}
	}
}
