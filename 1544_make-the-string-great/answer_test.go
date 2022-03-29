/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/29 15:19
 */

package _make_the_string_great

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, makeGood)

	testCase.SetAndRun("leEeetcode", "leetcode")
	testCase.SetAndRun("abBAcC", "")
	testCase.SetAndRun("s", "s")
	testCase.SetAndRun("aA", "")
}

func FuzzAnswer(f *testing.F) {
	f.Add("aA")
	f.Fuzz(func(t *testing.T, s string) {
		makeGood(s)
	})
}
