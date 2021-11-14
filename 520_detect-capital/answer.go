/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/detect-capital/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/14 8:35 下午
 */

package _detect_capital

func detectCapitalUse(word string) bool {
	if len(word) > 1 {
		if isCapital(word[0]) && isCapital(word[1]) {
			for _, b := range word[2:] {
				if !isCapital(byte(b)) {
					return false
				}
			}
		} else if (isCapital(word[0]) && !isCapital(word[1])) || (!isCapital(word[0])) {
			for _, b := range word[1:] {
				if isCapital(byte(b)) {
					return false
				}
			}
		}
	}

	return true
}

func isCapital(b byte) bool {
	return b >= 'A' && b <= 'Z'
}
