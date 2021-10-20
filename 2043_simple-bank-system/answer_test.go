/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/20 2:13 下午
 */

package _simple_bank_system

import "testing"

func TestAnswer(t *testing.T) {
	bank := Constructor([]int64{10, 100, 20, 50, 30})

	t.Logf("%t", bank.Withdraw(3, 10))              // true
	t.Logf("%t", bank.Transfer(5, 1, 20)) // true
	t.Logf("%t", bank.Deposit(5, 20))     	      // true
	t.Logf("%t", bank.Transfer(3, 4, 15)) // false
	t.Logf("%t", bank.Withdraw(10, 50))   		  // false
}
