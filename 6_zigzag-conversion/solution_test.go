/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/6/9 10:26
 */

package _zigzag_conversion

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestSolution(t *testing.T) {
	type input struct {
		s string
		numRows int
	}

	test := utils.NewTestCase(t, func(i input) string {
		return convert(i.s, i.numRows)
	})

	test.SetAndRun(input{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR")
	test.SetAndRun(input{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI")
	test.SetAndRun(input{"A", 1}, "A")
	test.SetAndRun(input{"AB", 1}, "AB")
	test.SetAndRun(input{"ABC", 1}, "ABC")
}
