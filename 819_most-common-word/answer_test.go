/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/17 13:19
 */

package _most_common_word

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	type input struct {
		paragraph string
		banned    []string
	}

	testCase := utils.NewTestCase(t, func(i input) string {
		return mostCommonWord(i.paragraph, i.banned)
	})

	testCase.SetAndRun(input{"Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}}, "ball")
	testCase.SetAndRun(input{"Bob bob BOB, bbb aaa ccc", []string{"aaa"}}, "bob")
	testCase.SetAndRun(input{"a, a, a, a, b,b,b,c, c", []string{"a"}}, "b")
	testCase.SetAndRun(input{"a, a, a, a,          b,b,b,c, c", []string{"a"}}, "b")
	testCase.SetAndRun(input{"a, a, a, a,  ?????        b,b,b,c, c", []string{"a"}}, "b")
	testCase.SetAndRun(input{"Bob", []string{}}, "bob")
}
