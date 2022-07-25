/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/7/25 10:38
 */

package _bulls_and_cows

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	tests := utils.NewTestCase(t, func(i [2]string) string {
		return getHint(i[0], i[1])
	})

	tests.SetAndRun([2]string{"1807", "7810"}, "1A3B")
	tests.SetAndRun([2]string{"1123", "0111"}, "1A1B")
	tests.SetAndRun([2]string{"1122", "1222"}, "3A0B")
}
