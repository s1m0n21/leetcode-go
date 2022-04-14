/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/14 12:59
 */

package _richest_customer_wealth

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, maximumWealth)

	testCase.SetAndRun([][]int{{1, 2, 3}, {3, 2, 1}}, 6)
	testCase.SetAndRun([][]int{{1, 5}, {7, 3}, {3, 5}}, 10)
	testCase.SetAndRun([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, 17)
}
