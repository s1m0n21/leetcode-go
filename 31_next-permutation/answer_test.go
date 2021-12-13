/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/12/13 2:19 PM
 */

package _next_permutation

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1}, []int{1}},
	}

	for i, test := range tests {
		c := append([]int(nil), test.input...)
		if nextPermutation(c); !reflect.DeepEqual(c, test.expect) {
			t.Errorf("%d: input = %+v, expect= %+v, actual = %+v", i, test.input, test.expect, c)
		}
	}
}
