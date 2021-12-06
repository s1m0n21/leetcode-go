/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/ransom-note/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/12/6 5:16 PM
 */

package _ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	var record [26]int
	for _, b := range magazine {
		record[b-'a']++
	}

	for _, b := range ransomNote {
		if record[b-'a'] == 0 {
			return false
		}
		record[b-'a']--
	}

	return true
}
