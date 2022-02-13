/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/2/13 4:19 PM
 */

package _pascals_triangle_ii

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  int
		expect []int
	}{
		{3, []int{1, 3, 3, 1}},
		{0, []int{1}},
		{1, []int{1, 1}},
	}

	for i, test := range tests {
		if actual := getRow(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %d, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
