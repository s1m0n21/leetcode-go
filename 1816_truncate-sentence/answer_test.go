/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/12/6 5:05 PM
 */

package _truncate_sentence

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		s string
		k int
	}

	tests := []struct {
		input  input
		expect string
	}{
		{input{"Hello how are you Contestant", 4}, "Hello how are you"},
		{input{"What is the solution to this problem", 4}, "What is the solution"},
		{input{"chopper is not a tanuki", 5}, "chopper is not a tanuki"},
	}

	for i, test := range tests {
		if actual := truncateSentence(test.input.s, test.input.k); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %s, actual = %s", i, test.input, test.expect, actual)
		}
	}
}
