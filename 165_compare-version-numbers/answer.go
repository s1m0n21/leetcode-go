/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/compare-version-numbers/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/1 8:59 上午
 */

package _compare_version_numbers

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	ver1 := strings.Split(version1, ".")
	ver2 := strings.Split(version2, ".")

	for len(ver1) > 0 || len(ver2) > 0 {
		if len(ver1) > 0 && len(ver2) > 0 {
			v1, _ := strconv.Atoi(ver1[0])
			v2, _ := strconv.Atoi(ver2[0])

			if v1 > v2 {
				return 1
			} else if v2 > v1 {
				return -1
			}

			ver1 = ver1[1:]
			ver2 = ver2[1:]
		} else if len(ver1) > 0 {
			v, _ := strconv.Atoi(ver1[0])
			if v > 0 {
				return 1
			}
			ver1 = ver1[1:]
		} else {
			v, _ := strconv.Atoi(ver2[0])
			if v > 0 {
				return -1
			}
			ver2 = ver2[1:]
		}
	}

	return 0
}
