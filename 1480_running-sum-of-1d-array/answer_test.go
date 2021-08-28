/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/28 10:09 下午
 */

package _running_sum_of_1d_array

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}

	for _, test := range tests {
		c := utils.CopySlice(test.input)
		if actual := runningSum(c); !reflect.DeepEqual(test.expect, actual) {
			t.Errorf("input = %#v, expect = %#v, actual = %#v", test.input, test.expect, actual)
		}
	}
}
