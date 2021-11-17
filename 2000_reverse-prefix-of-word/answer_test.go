/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/17 4:32 下午
 */

package _reverse_prefix_of_word

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		word string
		ch   byte
	}

	tests := []struct {
		input  input
		expect string
	}{
		{input{"abcdefd", 'd'}, "dcbaefd"},
		{input{"xyxzxe", 'z'}, "zxyxxe"},
		{input{"abcd", 'z'}, "abcd"},
		{input{"abcdefg", 'g'}, "gfedcba"},
	}

	for i, test := range tests {
		if actual := reversePrefix(test.input.word, test.input.ch); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %s, actual = %s", i, test.input, test.expect, actual)
		}
	}
}
