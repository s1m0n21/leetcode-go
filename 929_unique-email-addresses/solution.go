/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/unique-email-addresses/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/6/1 09:38
 */

package _unique_email_addresses

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	addrs := make(map[string]int)
	for _, addr := range emails {
		s := strings.Split(addr, "@")
		addrs[fmt.Sprintf("%s@%s", s[len(s)-1], strings.ReplaceAll(strings.Split(s[0], "+")[0], ".", ""))]++
	}
	return len(addrs)
}
