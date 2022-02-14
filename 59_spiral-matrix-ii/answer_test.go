/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/2/14 10:00 AM
 */

package _spiral_matrix_ii

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  int
		expect [][]int
	}{
		{3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{1, [][]int{{1}}},
	}

	for i, test := range tests {
		if actual := generateMatrix(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %d, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
