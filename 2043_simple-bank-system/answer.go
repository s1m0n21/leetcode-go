/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/simple-bank-system/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/20 2:03 下午
 */

package _simple_bank_system

type Bank struct {
	balances []int64
	n int
}

func Constructor(balance []int64) Bank {
	return Bank{
		balances: balance,
		n: len(balance),
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.validAccount(account1) || !this.validAccount(account2) || this.balances[account1-1] < money {
		return false
	}

	this.balances[account1-1] -= money
	this.balances[account2-1] += money

	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.validAccount(account) {
		return false
	}

	this.balances[account-1] += money

	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.validAccount(account) || this.balances[account-1] < money {
		return false
	}

	this.balances[account-1] -= money

	return true
}

func (this *Bank) validAccount(account int) bool {
	return account >= 1 && account <= this.n
}
