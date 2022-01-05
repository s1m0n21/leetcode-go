/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/restore-ip-addresses/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/5 12:37 PM
 */

package _restore_ip_addresses

import "strconv"

func restoreIpAddresses(s string) []string {
	var ans []string
	var dfs func(st, idx int)
	var seg = make([]int, 4)

	dfs = func(st, idx int) {
		if idx == 4 && st == len(s) {
			b := make([]byte, 0)
			for i := 0; i < 4; i++ {
				b = append(b, []byte(strconv.Itoa(seg[i]))...)
				if i != 3 {
					b = append(b, '.')
				}
			}
			ans = append(ans, string(b))
		}

		if idx == 4 || st == len(s) {
			return
		}

		if s[st] == '0' {
			seg[idx] = 0
			dfs(st+1, idx+1)
		}

		addr := 0
		for i := st; i < len(s); i++ {
			addr = addr*10 + int(s[i]-'0')
			if addr > 0 && addr <= 255 {
				seg[idx] = addr
				dfs(i+1, idx+1)
			} else {
				break
			}
		}
	}

	dfs(0, 0)

	return ans
}
