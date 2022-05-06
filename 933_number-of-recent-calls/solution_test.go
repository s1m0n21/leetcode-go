/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/6 09:34
 */

package _number_of_recent_calls

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	rc := Constructor()
	testCase := utils.NewTestCase(t, rc.Ping)

	testCase.SetAndRun(1, 1)
	testCase.SetAndRun(100, 2)
	testCase.SetAndRun(3001, 3)
	testCase.SetAndRun(3002, 3)
}
