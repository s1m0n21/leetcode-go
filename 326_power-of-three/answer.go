/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/power-of-three/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/31 9:31 上午
 */

 package _power_of_three

func isPowerOfThree(n int) bool {
  return n > 0 && 1162261467 % n == 0
}

