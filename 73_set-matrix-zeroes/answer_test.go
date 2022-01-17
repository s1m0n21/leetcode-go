/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/17 3:32 PM
 */

package _set_matrix_zeroes

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  [][]int
		expect [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
	}

	for i, test := range tests {
		var c [][]int
		for _, r := range test.input {
			c = append(c, r)
		}
		if setZeroes(c); !reflect.DeepEqual(c, test.expect) {
			t.Errorf("%d: input = %+v, exepct = %+v, actual = %+v", i, test.input, test.expect, c)
		}
	}
}
