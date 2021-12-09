/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/12/9 9:42 AM
 */

package _single_number_iii

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{1, 2, 1, 3, 2, 5}, []int{3, 5}},
		{[]int{-1, 0}, []int{-1, 0}},
		{[]int{1, 0}, []int{1, 0}},
	}

	for i, test := range tests {
		if actual := singleNumber(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
