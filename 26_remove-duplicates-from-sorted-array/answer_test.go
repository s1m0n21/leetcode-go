/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/25 4:07 下午
 */

package _remove_duplicates_from_sorted_array

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1,1,2}, 2},
		{[]int{0,0,1,1,1,2,2,3,3,4}, 5},
	}

	for _, test := range tests {
		c := utils.CopySlice(test.input)
		if actual := removeDuplicates(c);
			actual != test.expect || utils.HasDuplicate(c[:actual]) {
			t.Errorf("input = %+v, expect = %+v, actual = %d", test.input, test.expect, actual)
		}
	}
}
