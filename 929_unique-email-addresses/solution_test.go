/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/6/1 09:44
 */

package _unique_email_addresses

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	test := utils.NewTestCase(t, numUniqueEmails)

	test.SetAndRun([]string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}, 2)
	test.SetAndRun([]string{"a@leetcode.com","b@leetcode.com","c@leetcode.com"}, 3)
}
