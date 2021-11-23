/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/buddy-strings/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/23 10:30 上午
 */

package _buddy_strings

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	i, j := -1, -1
	chars := [26]int{}
	for k := 0; k < len(s); k++ {
		chars[s[k]-'a']++
		if s[k] != goal[k] {
			if i == -1 {
				i = k
			} else if j == -1 {
				j = k
			} else {
				return false
			}
		}
	}

	if i != -1 && j != -1 {
		return s[i] == goal[j] && s[j] == goal[i]
	}

	if i == -1 && j == -1 {
		for _, count := range chars {
			if count > 1 {
				return true
			}
		}
	}

	return false
}
