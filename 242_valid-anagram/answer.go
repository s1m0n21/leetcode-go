/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/valid-anagram/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/2 1:46 上午
 */

package _valid_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var need = [26]int{}
	for i := 0; i < len(s); i++ {
		need[s[i]-'a']++
	}

	var target = [26]int{}
	for i := 0; i < len(t); i++ {
		target[t[i]-'a']++
	}

	return target == need
}
