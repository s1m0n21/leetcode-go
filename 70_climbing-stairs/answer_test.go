/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/12 3:06 下午
 */

package _climbing_stairs

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  int
		expect int
	}{
		{2, 2},
		{3, 3},
	}

	for i, test := range tests {
		if actual := climbStairs(test.input); actual != test.expect {
			t.Errorf("%d: input = %d, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
