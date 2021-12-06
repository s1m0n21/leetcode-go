/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/12/6 5:20 PM
 */

package _ransom_note

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		ransomNote, magazine string
	}

	tests := []struct {
		input  input
		expect bool
	}{
		{input{"a", "b"}, false},
		{input{"aa", "ab"}, false},
		{input{"aa", "aab"}, true},
	}

	for i, test := range tests {
		if actual := canConstruct(test.input.ransomNote, test.input.magazine); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %t, actual = %t", i, test.input, test.expect, actual)
		}
	}
}
