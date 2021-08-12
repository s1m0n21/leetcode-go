/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/12 12:56 下午
 */

package _combination_sum_iii

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		k, n int
	}

	tests := []struct {
		input  input
		expect [][]int
	}{
		{input{3, 7}, [][]int{{1, 2, 4}}},
		{input{3, 9}, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
	}

	for _, test := range tests {
		if actual := combinationSum3(test.input.k, test.input.n); !reflect.DeepEqual(test.expect, actual) {
			t.Errorf("input = %+v, expect = %+v, actual = %+v", test.input, test.expect, actual)
		}
	}
}
