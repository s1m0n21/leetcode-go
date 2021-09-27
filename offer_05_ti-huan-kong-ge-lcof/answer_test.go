/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/9/27 2:33 下午
 */

package offer_05_ti_huan_kong_ge_lcof

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"We are happy.", "We%20are%20happy."},
		{"包含 中文 呢", "包含%20中文%20呢"},
	}

	for _, test := range tests {
		if actual := replaceSpace(test.input); actual != test.expect {
			t.Errorf("input = %s, expect = %s, actual = %s", test.input, test.expect, actual)
		}
	}
}
