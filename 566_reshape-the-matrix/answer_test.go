/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/13 12:56 PM
 */

package _reshape_the_matrix

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		mat  [][]int
		r, c int
	}

	tests := []struct {
		input  input
		expect [][]int
	}{
		{input{[][]int{{1, 2}, {3, 4}}, 1, 4}, [][]int{{1, 2, 3, 4}}},
		{input{[][]int{{1, 2}, {3, 4}}, 2, 4}, [][]int{{1, 2}, {3, 4}}},
	}

	for i, test := range tests {
		if actual := matrixReshape(test.input.mat, test.input.r, test.input.c); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
