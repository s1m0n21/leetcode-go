/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/30 1:17 上午
 */

package _find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	if len(s) == 0 || len(s) < len(p) {
		return nil
	}

	var out = make([]int, 0)
	var appeared [26]int
	var target [26]int
	for i := 0; i < len(p); i++ {
		target[p[i]-'a']++
	}

	left, right := 0, len(p)-1
	for i := left; i <= right; i++ {
		appeared[s[i]-'a']++
	}

	for right < len(s) {
		if appeared == target {
			out = append(out, left)
		}

		appeared[s[left]-'a']--

		right++
		left++

		if right < len(s) {
			appeared[s[right]-'a']++
		}
	}

	return out
}
