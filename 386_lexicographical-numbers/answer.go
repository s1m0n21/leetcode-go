/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/lexicographical-numbers/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/18 14:08
 */

package _lexicographical_numbers

func lexicalOrder(n int) []int {
	ret := make([]int, n)
	num := 1
	for i := range ret {
		ret[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}
	return ret
}
