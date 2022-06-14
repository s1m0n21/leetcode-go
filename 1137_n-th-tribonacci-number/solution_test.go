/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/6/14 15:04
 */

package _n_th_tribonacci_number

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	test := utils.NewTestCase(t, tribonacci)

	test.SetAndRun(4, 4)
	test.SetAndRun(3, 2)
	test.SetAndRun(0, 0)
	test.SetAndRun(1, 1)
	test.SetAndRun(2, 1)
	test.SetAndRun(25, 1389537)
}
