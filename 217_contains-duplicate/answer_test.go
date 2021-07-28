/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/17 11:34 下午
 */

package _contains_duplicate

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect bool
	}{
		{[]int{1, 2, 3}, false},
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 1, 1, 2, 3}, true},
		{[]int{}, false},
		{[]int{-1, 1, 2, 3}, false},
	}

	for _, test := range tests {
		if actual := containsDuplicate(test.input); actual != test.expect {
			t.Errorf("input = %+v, expect = %t, actual = %t", test.input, test.expect, actual)
		}
	}
}
