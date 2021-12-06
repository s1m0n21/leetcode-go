/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/consecutive-characters/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/12/6 5:31 PM
 */

package _consecutive_characters

func maxPower(s string) int {
	var max, tmp int
	for i, b := range s {
		if i == 0 || s[i-1] != byte(b) {
			if tmp > max {
				max = tmp
			}
			tmp = 1
		} else {
			tmp++
		}
	}

	if tmp > max {
		max = tmp
	}

	return max
}
