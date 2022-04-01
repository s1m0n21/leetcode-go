/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/count-common-words-with-one-occurrence/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/1 17:38
 */

package _count_common_words_with_one_occurrence

func countWords(words1 []string, words2 []string) int {
	record1 := make(map[string]int)
	for _, s := range words1 {
		record1[s]++
	}
	record2 := make(map[string]int)
	for _, s := range words2 {
		record2[s]++
	}

	var ret int
	for _, s := range words1 {
		r1 := record1[s]
		r2 := record2[s]
		if r1 > 0 && r2 > 0 && r1+r2 == 2 {
			ret++
		}
	}

	return ret
}
