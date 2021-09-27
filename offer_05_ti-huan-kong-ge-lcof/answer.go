/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/27 2:28 下午
 */

package offer_05_ti_huan_kong_ge_lcof

func replaceSpace(s string) string {
	if len(s) == 0 {
		return s
	}

	var space = []byte("%20")
	var b []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			b = append(b, space...)
		} else {
			b = append(b, s[i])
		}
	}

	return string(b)
}
