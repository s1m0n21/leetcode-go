/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/compress-string-lcci/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/21 3:32 下午
 */

package interview_01_06_compress_string_lcci

import "strconv"

func compressString(s string) string {
	if len(s) == 0 {
		return ""
	}

	var ans []byte
	l, r := 0, 0
	for ; r < len(s); r++ {
		if s[r] != s[l] {
			ans = append(ans, s[l])
			ans = append(ans, []byte(strconv.Itoa(r-l))...)
			l = r
		}
	}
	ans = append(ans, s[l])
	ans = append(ans, []byte(strconv.Itoa(r-l))...)

	if len(ans) >= len(s) {
		return s
	}

	return string(ans)
}
