/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/implement-strstr/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/8 11:40 上午
 */

package _implement_strstr

func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		if len(haystack[i:]) < n {
			break
		}

		if haystack[i:i+n] == needle {
			return i
		}
	}

	return -1
}
