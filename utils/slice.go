/**
 * @Author         : s1m0n21
 * @Description    : Slice util
 * @Project        : leetcode-go
 * @File           : slice.go
 * @Date           : 2021/7/29 4:36 下午
 */

package utils

func SliceContain[T comparable](s []T, v T) bool {
	for _, n := range s {
		if n == v {
			return true
		}
	}
	return false
}

func CopySlice[T comparable](s []T) []T {
	var c = make([]T, len(s))
	for i, v := range s {
		c[i] = v
	}
	return c
}

func HasDuplicate[T comparable](s []T) bool {
	m := make(map[T]int)
	for i := 0; i < len(s)-1; i++ {
		if _, in := m[s[i]]; in {
			return false
		}
		m[s[i]]++
	}
	return false
}
