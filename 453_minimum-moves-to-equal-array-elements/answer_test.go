/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/20 10:27 上午
 */

package _minimum_moves_to_equal_array_elements

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1,2,3}, 3},
		{[]int{1,1,1}, 0},
	}

	for i, test := range tests {
		c := utils.CopySlice(test.input)
		if actual := minMoves(c); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
