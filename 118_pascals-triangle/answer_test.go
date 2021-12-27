/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/12/27 12:53 PM
 */

package _pascals_triangle

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  int
		expect [][]int
	}{
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{1, [][]int{{1}}},
	}

	for i, test := range tests {
		if actual := generate(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %d, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
