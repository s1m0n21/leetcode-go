/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/4 12:56 PM
 */

package _gas_station

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		gas, cost []int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}}, 3},
		{input{[]int{2, 3, 4}, []int{3, 4, 3}}, -1},
	}

	for i, test := range tests {
		if actual := canCompleteCircuit(test.input.gas, test.input.cost); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
