/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/17 10:26
 */

package _verifying_an_alien_dictionary

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	type input struct {
		words []string
		order string
	}

	testCase := utils.NewTestCase(t, func(i input) bool {
		return isAlienSorted(i.words, i.order)
	})

	testCase.SetAndRun(input{[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"}, false)
	testCase.SetAndRun(input{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"}, true)
	testCase.SetAndRun(input{[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"}, false)
}
