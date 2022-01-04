/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/4 12:17 PM
 */

package _find_right_interval

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  [][]int
		expect []int
	}{
		{[][]int{{1, 2}}, []int{-1}},
		{[][]int{{3, 4}, {2, 3}, {1, 2}}, []int{-1, 0, 1}},
		{[][]int{{1, 4}, {2, 3}, {3, 4}}, []int{-1, 2, -1}},
		{[][]int{{1, 1}, {3, 4}}, []int{0, -1}},
	}

	for i, test := range tests {
		if actual := findRightInterval(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
