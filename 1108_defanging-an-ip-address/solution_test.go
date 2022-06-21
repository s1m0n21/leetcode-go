/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/6/21 09:52
 */

package _defanging_an_ip_address

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestSolution(t *testing.T) {
	test := utils.NewTestCase(t, defangIPaddr)

	test.SetAndRun("1.1.1.1", "1[.]1[.]1[.]1")
	test.SetAndRun("255.100.50.0", "255[.]100[.]50[.]0")
}
