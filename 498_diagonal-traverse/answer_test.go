/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/15 1:03 上午
 */

package _diagonal_traverse

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  [][]int
		expect []int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[]int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			[][]int{
				{1, 2},
				{3, 4},
			},
			[]int{1, 2, 3, 4},
		},
		{
			[][]int{
				{1},
			},
			[]int{1},
		},
	}

	for i, test := range tests {
		if actual := findDiagonalOrder(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
