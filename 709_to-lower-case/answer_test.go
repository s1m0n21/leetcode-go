/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/9/27 2:21 下午
 */

package _to_lower_case

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"Hello", "hello"},
		{"here", "here"},
		{"LOVELY", "lovely"},
		{"Hello,WOrld!", "hello,world!"},
	}

	for _, test := range tests {
		if actual := toLowerCase(test.input); actual != test.expect {
			t.Errorf("input = %s, expect = %s, actual = %s", test.input, test.expect, actual)
		}
	}
}
