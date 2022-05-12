/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/12 09:55
 */

package _delete_columns_to_make_sorted

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	testCase := utils.NewTestCase(t, minDeletionSize)

	testCase.SetAndRun([]string{"cba", "daf", "ghi"}, 1)
	testCase.SetAndRun([]string{"a", "b"}, 0)
	testCase.SetAndRun([]string{"zyx","wvu","tsr"}, 3)
}
