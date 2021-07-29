/**
 * @Author         : s1m0n21
 * @Description    : Test Answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/25 2:58 下午
 */

package _remove_element

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		nums []int
		val  int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{[]int{3, 2, 2, 3}, 3}, 2},
		{input{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, 5},
	}

	for _, test := range tests {
		c := utils.CopySlice(test.input.nums)
		if actual := removeElement(c, test.input.val);
		actual != test.expect || utils.SliceContain(c[:actual], test.input.val) {
			t.Errorf("input = %+v, expect = %+v, actual = %d", test.input, test.expect, actual)
		}
	}
}
