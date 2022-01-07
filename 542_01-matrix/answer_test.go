/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/7 12:31 PM
 */

package _01_matrix

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  [][]int
		expect [][]int
	}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
		{[][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}},
	}

	for i, test := range tests {
		if actual := updateMatrix(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %+v, exepct = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
