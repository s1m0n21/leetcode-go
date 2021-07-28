/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/sort-characters-by-frequency/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/4 12:41 上午
 */

package _sort_characters_by_frequency

import "sort"

type pair struct {
	Key   byte
	Value int
}

func frequencySort(s string) string {
	var freq = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	var pairs = make([]pair, 0)
	for k, v := range freq {
		pairs = append(pairs, pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	var out = make([]byte, 0)
	for _, p := range pairs {
		for i := 0; i < p.Value; i++ {
			out = append(out, p.Key)
		}
	}

	return string(out)
}
