/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/2 10:58
 */

package _merge_strings_alternately

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, func(i [2]string) string {
		return mergeAlternately(i[0], i[1])
	})

	testCase.SetAndRun([2]string{"abc", "pqr"}, "apbqcr")
	testCase.SetAndRun([2]string{"ab", "pqrs"}, "apbqrs")
	testCase.SetAndRun([2]string{"abcd", "pq"}, "apbqcd")
	testCase.SetAndRun([2]string{"a", "b"}, "ab")
}
