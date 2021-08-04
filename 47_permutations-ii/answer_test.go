/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/4 5:05 下午
 */

package _permutations_ii

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect [][]int
	}{
		{[]int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}

	for _, test := range tests {
		if actual := permuteUnique(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("input = %v, expect = %v, actual = %v", test.input, test.expect, actual)
		}
	}
}
