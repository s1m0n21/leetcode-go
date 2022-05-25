/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/25 18:03
 */

package _unique_substrings_in_wraparound_string

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	test := utils.NewTestCase(t, findSubstringInWraproundString)

	test.SetAndRun("a", 1)
	test.SetAndRun("cac", 2)
	test.SetAndRun("zab", 6)
	test.SetAndRun("zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd", 1157)
}
