/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/4 10:02 下午
 */

package _longest_substring_without_repeating_characters

//func lengthOfLongestSubstring(s string) int {
//	if s == "" {
//		return 0
//	}
//
//	freq := map[byte]int{}
//	left, right := 0, -1
//	out := 0
//
//	for left < len(s) {
//		if right+1 < len(s) && freq[s[right+1]] == 0 {
//			right++
//			freq[s[right]]++
//		} else {
//			freq[s[left]]--
//			left++
//		}
//
//		out = max(out, right-left+1)
//	}
//
//	return out
//}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	var appeared [256]int
	for i := 0; i < len(appeared); i++ {
		appeared[i] = -1
	}

	left, right := 0, 0
	out := 0

	for right < len(s) {
		if i := appeared[s[right]-'a']; i >= left {
			left = i + 1
		}

		out = max(out, right-left+1)

		appeared[s[right]-'a'] = right
		right++
	}

	return out
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
