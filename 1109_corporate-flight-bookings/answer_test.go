/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/31 12:29 上午
 */

package _corporate_flight_bookings

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		bookings [][]int
		n        int
	}

	tests := []struct {
		input  input
		expect []int
	}{
		{input{[][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5}, []int{10, 55, 45, 25, 25}},
		{input{[][]int{{1, 2, 10}, {2, 2, 15}}, 2}, []int{10, 25}},
	}

	for _, test := range tests {
		if actual := corpFlightBookings(test.input.bookings, test.input.n); !reflect.DeepEqual(test.expect, actual) {
			t.Errorf("input = %+v, expect = %+v, actual = %+v", test.input, test.expect, actual)
		}
	}
}
