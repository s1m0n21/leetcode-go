/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/first-unique-character-in-a-string/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/17 10:12 AM
 */

package _first_unique_character_in_a_string

func firstUniqChar(s string) int {
	var record [26]int

	for _, b := range s {
		record[b-'a']++
	}
	for i, b := range s {
		if record[b-'a'] == 1 {
			return i
		}
	}

	return -1
}
