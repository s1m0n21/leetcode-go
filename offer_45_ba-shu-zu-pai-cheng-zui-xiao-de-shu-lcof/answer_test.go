/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/7 5:56 下午
 */

package offer_45_ba_shu_zu_pai_cheng_zui_xiao_de_shu_lcof

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect string
	}{
		{[]int{10, 2}, "102"},
		{[]int{3,30,34,5,9}, "3033459"},
		{[]int{1,2,3,1}, "1123"},
		{[]int{1,2,3,4,5,6,7,8,9,0}, "0123456789"},
	}

	for _, test := range tests {
		c := utils.CopySlice(test.input)
		if actual := minNumber(c); actual != test.expect {
			t.Errorf("input = %#v, expect = %s, actual = %s", test.input, test.expect, actual)
		}
	}
}
