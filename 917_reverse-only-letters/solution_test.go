/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/6/7 09:39
 */

package _reverse_only_letters

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	test := utils.NewTestCase(t, reverseOnlyLetters)

	test.SetAndRun("ab-cd", "dc-ba")
	test.SetAndRun("a-bC-dEf-ghIj", "j-Ih-gfE-dCba")
	test.SetAndRun("Test1ng-Leet=code-Q!", "Qedo1ct-eeLg=ntse-T!")
	test.SetAndRun("!@!@!@!@", "!@!@!@!@")
}
