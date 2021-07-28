/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/simplify-path/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/23 12:11 上午
 */

package _simplify_path

import (
	"strings"
)

func simplifyPath(path string) string {
	var stack = make([]string, 0)
	for _, s := range strings.Split(path, "/") {
		if s == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if s == "." {
			continue
		} else if len(s) > 0 {
			stack = append(stack, s)
		}
	}

	res := ""
	for _, s := range stack {
		res = res + "/" + s
	}

	if res == "" {
		res = "/"
	}

	return res
}
