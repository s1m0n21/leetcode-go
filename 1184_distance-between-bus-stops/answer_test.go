/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/17 3:50 下午
 */

package _distance_between_bus_stops

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		distance           []int
		start, destination int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{[]int{1,2,3,4}, 0, 1}, 1},
		{input{[]int{1,2,3,4}, 0, 2}, 3},
		{input{[]int{1,2,3,4}, 0, 3}, 4},
		{input{[]int{7,10,1,12,11,14,5, 0}, 7, 2}, 17},
	}

	for i, test := range tests {
		if actual := distanceBetweenBusStops(test.input.distance, test.input.start, test.input.destination); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
