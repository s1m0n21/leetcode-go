/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/7/25 10:56
 */

package _backspace_string_compare

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	tests := utils.NewTestCase(t, func(i [2]string) bool {
		return backspaceCompare(i[0], i[1])
	})

	tests.SetAndRun([2]string{"ab#c", "ad#c"}, true)
	tests.SetAndRun([2]string{"ab##", "c#d#"}, true)
	tests.SetAndRun([2]string{"a#c", "b"}, false)
	tests.SetAndRun([2]string{"#", "b#"}, true)
}
