/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/2/11 10:56 AM
 */

package _merge_intervals

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  [][]int
		expect [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{0, 9}, {4, 5}}, [][]int{{0, 9}}},
		{[][]int{{1, 2}, {3, 5}, {4, 7}}, [][]int{{1, 2}, {3, 7}}},
		{[][]int{{1, 2}, {2, 2}, {4, 7}}, [][]int{{1, 2}, {4, 7}}},
	}

	for i, test := range tests {
		if actual := merge(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
