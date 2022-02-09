/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/longest-happy-string/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/2/8 10:13 AM
 */

package _longest_happy_string

import "sort"

func longestDiverseString(a int, b int, c int) string {
	var ans []byte

	cnt := []struct {
		c  int
		ch byte
	}{
		{a, 'a'},
		{b, 'b'},
		{c, 'c'},
	}

	for {
		sort.Slice(cnt, func(i, j int) bool {
			return cnt[i].c > cnt[j].c
		})
		hasNext := false
		for i, p := range cnt {
			if p.c <= 0 {
				break
			}
			n := len(ans)
			if n >= 2 && ans[n-2] == p.ch && ans[n-1] == p.ch {
				continue
			}
			hasNext = true
			ans = append(ans, p.ch)
			cnt[i].c--

			break
		}
		if !hasNext {
			return string(ans)
		}
	}
}
