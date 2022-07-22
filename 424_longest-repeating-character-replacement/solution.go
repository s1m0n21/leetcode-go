/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/longest-repeating-character-replacement/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/7/22 15:26
 */

package _longest_repeating_character_replacement

func characterReplacement(s string, k int) int {
	cnt := [26]int{}
	maxCnt, left, right := 0, 0, 0
	for ; right < len(s); right++ {
		cnt[s[right]-'A']++
		maxCnt = max(cnt[s[right]-'A'], maxCnt)
		if right-left+1-maxCnt > k {
			cnt[s[left]-'A']--
			left++
		}
	}

	return len(s) - left
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
