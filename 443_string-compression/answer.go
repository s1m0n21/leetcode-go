/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/string-compression/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/1 3:22 下午
 */

package _string_compression

import "strconv"

func compress(chars []byte) int {
	length := len(chars)
	if length == 0 || length == 1 {
		return length
	}

	j, k := 0, 0
	for i := 0; i < length; i++ {
		if i+1 == length || chars[i+1] != chars[i] {
			chars[j] = chars[k]
			j++

			if i > k {
				n := []byte(strconv.Itoa(i + 1 - k))
				for _, b := range n {
					chars[j] = b
					j++
				}
			}

			k = i + 1
		}
	}

	return j
}
