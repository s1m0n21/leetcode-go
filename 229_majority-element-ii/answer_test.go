/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/22 5:08 下午
 */

package _majority_element_ii

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{3, 2, 3}, []int{3}},
		{[]int{1}, []int{1}},
		{[]int{1, 1, 1, 3, 3, 2, 2, 2}, []int{1, 2}},
	}

	for i, test := range tests {
		if actual := majorityElement(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
