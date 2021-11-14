/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/14 8:44 下午
 */

package _detect_capital

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{"USA", true},
		{"Abc", true},
		{"hELLO", false},
		{"a", true},
		{"test", true},
		{"HeLlo", false},
		{"mL", false},
	}

	for i, test := range tests {
		if actual := detectCapitalUse(test.input); actual != test.expect {
			t.Errorf("%d: input = %s, expect = %t, actual = %t", i, test.input, test.expect, actual)
		}
	}
}
