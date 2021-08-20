/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/reverse-string-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/20 4:39 下午
 */

package _reverse_string_ii

func reverseStr(s string, k int) string {
	res := []byte(s)

	reverse := func(l, r int) {
		for l < r {
			res[l], res[r] = res[r], res[l]
			l++
			r--
		}
	}

	l, r := 0, k-1
	for r < len(s) {
		reverse(l, r)
		l = r + k + 1
		r = l + k - 1
	}

	if len(s) - l < k {
		reverse(l, len(s)-1)
	}

	return string(res)
}
