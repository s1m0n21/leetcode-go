/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/13 11:12 上午
 */

package _generate_parentheses

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  int
		expect []string
	}{
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{2, []string{"(())", "()()"}},
		{1, []string{"()"}},
	}

	for _, test := range tests {
		if actual := generateParenthesis(test.input); !reflect.DeepEqual(test.expect, actual) {
			t.Errorf("input = %d, expect = %#v, actual = %#v", test.input, test.expect, actual)
		}
	}
}
