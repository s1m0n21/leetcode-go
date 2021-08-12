/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/12 10:29 上午
 */

package _combination_sum_ii

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
		{input{[]int{10, 1, 2, 7, 6, 1, 5}, 8}, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{input{[]int{2, 5, 2, 1, 2}, 5}, [][]int{{1, 2, 2}, {5}}},
	}

	for _, test := range tests {
		if actual := combinationSum2(test.input.candidates, test.input.target); !reflect.DeepEqual(test.expect, actual) {
			t.Errorf("input = %+v, expect = %+v, actual = %+v", test.input, test.expect, actual)
		}
	}
}
