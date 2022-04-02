/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/2 16:35
 */

package _replace_all_s_to_avoid_consecutive_repeating_characters

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, modifyString)

	testCase.SetAndRun("?zs", "azs")
	testCase.SetAndRun("ubv?w", "ubvaw")
	testCase.SetAndRun("ubv??", "ubvab")
	testCase.SetAndRun("?", "a")
}
