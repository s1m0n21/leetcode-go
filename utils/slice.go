/**
 * @Author         : s1m0n21
 * @Description    : Slice util
 * @Project        : leetcode-go
 * @File           : outputSlice.go
 * @Date           : 2021/7/29 4:36 下午
 */

package utils

import "sort"

func SliceContain(s []int, v int) bool {
	for _, n := range s {
		if n == v {
			return true
		}
	}
	return false
}

func CopySlice(s []int) []int {
	var c = make([]int, len(s))
	for i, v := range s {
		c[i] = v
	}
	return c
}

func HasDuplicate(s []int) bool {
	sort.Ints(s)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}
