/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/string-to-url-lcci/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/18 17:21
 */

package interview_01_03_string_to_url_lcci

func replaceSpaces(s string, length int) string {
	var ret []byte

	for i := 0; i < length; i++ {
		if s[i] == ' ' {
			ret = append(ret, '%', '2', '0')
		} else {
			ret = append(ret, s[i])
		}
	}

	return string(ret)
}
