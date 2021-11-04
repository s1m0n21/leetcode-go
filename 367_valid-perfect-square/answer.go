/*
 * @Author       : s1m0n21
 * @Description  : Answer for https://leetcode-cn.com/problems/valid-perfect-square/
 * @Date         : 2021/11/04 12:06 PM
 */

package _valid_perfect_square

func isPerfectSquare(num int) bool {
	var a = 1
	for a*a < num {
		a++
	}
	return a*a == num
}
