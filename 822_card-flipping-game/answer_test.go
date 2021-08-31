/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/31 12:06 下午
 */

package _card_flipping_game

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		fronts, backs []int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{[]int{1, 2, 4, 4, 7}, []int{1, 3, 4, 1, 3}}, 2},
		{input{[]int{1, 1}, []int{1, 2}}, 2},
	}

	for _, test := range tests {
		f, b := utils.CopySlice(test.input.fronts), utils.CopySlice(test.input.backs)
		if actual := flipgame(f, b); actual != test.expect {
			t.Errorf("input = %+v, expect = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
