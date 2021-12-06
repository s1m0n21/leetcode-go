/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/12/6 4:46 PM
 */

package _reduce_array_size_to_the_half

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}, 2},
		{[]int{7, 7, 7, 7, 7, 7}, 1},
		{[]int{1, 9}, 1},
		{[]int{1000, 1000, 3, 7}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
	}

	for i, test := range tests {
		if actual := minSetSize(test.input); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
