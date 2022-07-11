/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/7/11 10:31
 */

package _implement_magic_dictionary

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	dict := Constructor()
	dict.BuildDict([]string{"hello", "leetcode"})

	tests := utils.NewTestCase(t, dict.Search)

	tests.SetAndRun("hello", false)
	tests.SetAndRun("hhllo", true)
	tests.SetAndRun("hell", false)
	tests.SetAndRun("leetcoded", false)
}
