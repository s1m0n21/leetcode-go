/**
 * @Author         : s1m0n21
 * @Description    : Answer for
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/12/6 4:40 PM
 */

package _reduce_array_size_to_the_half

import "sort"

func minSetSize(arr []int) int {
	freq := make(map[int]int)
	for _, n := range arr {
		freq[n]++
	}

	var s []int
	for _, n := range freq {
		s = append(s, n)
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})

	var ans, sum int
	for _, n := range s {
		if sum >= len(arr)/2 {
			break
		}
		sum += n
		ans++
	}

	return ans
}
