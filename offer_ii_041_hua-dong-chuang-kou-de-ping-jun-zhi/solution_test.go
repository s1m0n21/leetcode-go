/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/7/16 13:41
 */

package offer_ii_041_hua_dong_chuang_kou_de_ping_jun_zhi

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	avg := Constructor(3)

	tests := utils.NewTestCase(t, avg.Next)

	tests.SetAndRun(1, 1.0)
	tests.SetAndRun(10, 5.5)
	tests.SetAndRun(3, 4.666666666666667)
	tests.SetAndRun(5, 6.0)
}
