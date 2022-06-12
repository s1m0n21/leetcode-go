/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/6/12 15:44
 */

package _find_and_replace_pattern

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	type input struct {
		words   []string
		pattern string
	}

	test := utils.NewTestCase(t, func(i input) []string {
		return findAndReplacePattern(i.words, i.pattern)
	})
	test.ConfigSetOutputInAnyOrder()

	test.SetAndRun(input{[]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"}, []string{"mee", "aqq"})
	test.SetAndRun(input{[]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "aaa"}, []string{"ccc"})
}
