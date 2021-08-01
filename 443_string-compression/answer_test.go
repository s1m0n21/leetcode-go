/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/1 4:15 下午
 */

package _string_compression

import "testing"

func TestAnswer(t *testing.T) {
	type expect struct {
		length   int
		chars []byte
	}

	tests := []struct {
		input  []byte
		expect expect
	}{
		{[]byte("aaabbcc"), expect{6, []byte("a3b2c2")}},
		{[]byte("aabbccc"), expect{6, []byte("a2b2c3")}},
		{[]byte("a"), expect{1, []byte("a")}},
	}

	for _, test := range tests {
		actual := compress(test.input)
		if actual != test.expect.length || string(test.input[:actual]) != string(test.expect.chars) {
			t.Errorf("input = %s, expect = %+v, actual = %d", test.input, test.expect, actual)
		}
	}
}
