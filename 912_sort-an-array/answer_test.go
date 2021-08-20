/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/18 10:07 上午
 */

package _sort_an_array

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
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{0, 4, 2, 1, 3, 5}, []int{0, 1, 2, 3, 4, 5}},
		{[]int{5, 1, 1, 2, 0, 0}, []int{0, 0, 1, 1, 2, 5}},
	}

	for _, test := range tests {
		c := utils.CopySlice(test.input)
		if actual := sortArray(c); !reflect.DeepEqual(test.expect, actual) {
			t.Errorf("input = %+v, expect = %+v, actual = %+v", test.input, test.expect, actual)
		}
	}
}
