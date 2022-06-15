/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/6/15 09:20
 */

package _crawler_log_folder

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	test := utils.NewTestCase(t, minOperations)

	test.SetAndRun([]string{"d1/", "d2/", "../", "d21/", "./"}, 2)
	test.SetAndRun([]string{"d1/", "d2/", "./", "d3/", "../", "d31/"}, 3)
	test.SetAndRun([]string{"d1/", "../", "../", "../"}, 0)
}
