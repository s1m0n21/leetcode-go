/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/6 9:48 下午
 */

package _third_maximum_number

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 2, 3, 4}, 2},
		{[]int{1, 2, 2, 4, 2, 1}, 1},
		{[]int{1, 2}, 2},
		{[]int{-3, -2, -1, 0, 1, 2, 3, 4}, 2},
		{[]int{1, 2, -2147483648}, -2147483648},
	}

	for i, test := range tests {
		if actual := thirdMax(test.input); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
