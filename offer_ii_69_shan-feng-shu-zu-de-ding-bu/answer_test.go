/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/14 9:28 上午
 */

package offer_ii_69_shan_feng_shu_zu_de_ding_bu

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 3, 5, 4, 2}, 2},
		{[]int{0, 10, 5, 2}, 1},
		{[]int{3, 4, 5, 1}, 2},
		{[]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}, 2},
	}

	for i, test := range tests {
		if actual := peakIndexInMountainArray(test.input); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
