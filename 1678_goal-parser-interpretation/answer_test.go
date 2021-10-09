/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/9 9:34 上午
 */

package _goal_parser_interpretation

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"G()(al)", "Goal"},
		{"G()()()()(al)", "Gooooal"},
		{"(al)G(al)()()G", "alGalooG"},
	}

	for i, test := range tests {
		if actual := interpret(test.input); actual != test.expect {
			t.Errorf("%d: input = %s, expect = %s, actual = %s", i, test.input, test.expect, actual)
		}
	}
}
