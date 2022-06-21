/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/defanging-an-ip-address/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/6/21 09:48
 */

package _defanging_an_ip_address

func defangIPaddr(address string) string {
	var ret = make([]byte, 0, len(address)+6)
	for _, b := range address {
		if b == '.' {
			ret = append(ret, '[', '.', ']')
		} else {
			ret = append(ret, byte(b))
		}
	}

	return string(ret)
}
