/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/1 10:32
 */

package _summary_ranges

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, summaryRanges)

	testCase.SetAndRun([]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"})
	testCase.SetAndRun([]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"})
	testCase.SetAndRun(nil, nil)
	testCase.SetAndRun([]int{0}, []string{"0"})
}
