/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/11 10:28 下午
 */

package _combination_sum

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		candidates []int
		target     int
	}

	tests := []struct {
		input  input
		expect [][]int
	}{
		{input{[]int{2, 3, 6, 7}, 7}, [][]int{{2, 2, 3}, {7}}},
		{input{[]int{2, 3, 5}, 8}, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{input{[]int{2}, 1}, nil},
		{input{[]int{1}, 1}, [][]int{{1}}},
		{input{[]int{1}, 2}, [][]int{{1, 1}}},
	}

	for _, test := range tests {
		if actual := combinationSum(test.input.candidates, test.input.target); !reflect.DeepEqual(test.expect, actual) {
			t.Errorf("input = %+v, expect = %#v, actual = %#v", test.input, test.expect, actual)
		}
	}
}
