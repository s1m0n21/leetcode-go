/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/10 8:25 PM
 */

package _single_element_in_a_sorted_array

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2},
		{[]int{3, 3, 7, 7, 10, 11, 11}, 10},
	}

	for i, test := range tests {
		if actual := singleNonDuplicate(test.input); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
