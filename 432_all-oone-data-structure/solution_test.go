/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/7 15:52
 */

package _all_oone_data_structure

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestSolution(t *testing.T) {
	a := Constructor()
	testCase := utils.NewTestCase(t, func(i int) string {
		if i == 0 {
			return a.GetMinKey()
		}
		return a.GetMaxKey()
	})

	a.Inc("hello")
	a.Inc("hello")
	testCase.SetAndRun(1, "hello")
	testCase.SetAndRun(0, "hello")

	a.Inc("leet")
	testCase.SetAndRun(1, "hello")
	testCase.SetAndRun(0, "leet")

	a.Inc("leet")
	testCase.SetAndRun(1, "leet")
	testCase.SetAndRun(0, "hello")

	a.Dec("leet")
	testCase.SetAndRun(0, "leet")
	testCase.SetAndRun(1, "hello")

	a.Dec("leet")
	testCase.SetAndRun(0, "hello")
	testCase.SetAndRun(1, "hello")

	a.Inc("leet")
	testCase.SetAndRun(1, "hello")
	testCase.SetAndRun(0, "leet")

	a.Inc("hello")
	testCase.SetAndRun(1, "hello")
	testCase.SetAndRun(0, "leet")

	a.Inc("leet")
	testCase.SetAndRun(1, "hello")
	testCase.SetAndRun(0, "leet")

	a.Inc("code")
	testCase.SetAndRun(1, "hello")
	testCase.SetAndRun(0, "code")

	a.Dec("code")
	testCase.SetAndRun(1, "hello")
	testCase.SetAndRun(0, "leet")
}
