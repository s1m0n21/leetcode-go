/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/11 3:13 下午
 */

package _combinations

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		n, k int
	}

	tests := []struct {
		input  input
		expect [][]int
	}{
		{input{4, 2}, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{input{1, 1}, [][]int{{1}}},
	}

	for _, test := range tests {
		if actual := combine(test.input.n, test.input.k); !reflect.DeepEqual(test.expect, actual) {
			t.Errorf("input = %+v, expect = %+v, actual = %+v", test.input, test.expect, actual)
		}
	}
}
