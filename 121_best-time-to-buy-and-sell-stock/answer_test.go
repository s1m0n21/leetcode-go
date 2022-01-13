/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/13 8:07 PM
 */

package _best_time_to_buy_and_sell_stock

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{6, 5, 4, 3, 2, 1}, 0},
		{[]int{7, 4, 2, 1, 3, 8, 9, 10}, 9},
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 0, 5, 3, 6, 4}, 6},
		{[]int{1}, 0},
	}

	for i, test := range tests {
		if actual := maxProfit(test.input); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
