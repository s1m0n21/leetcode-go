/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/18 17:23
 */

package interview_01_03_string_to_url_lcci

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	type input struct {
		s      string
		length int
	}

	testCase := utils.NewTestCase(t, func(i input) string {
		return replaceSpaces(i.s, i.length)
	})

	testCase.SetAndRun(input{"Mr John Smith    ", 13}, "Mr%20John%20Smith")
	testCase.SetAndRun(input{"               ", 5}, "%20%20%20%20%20")
}
