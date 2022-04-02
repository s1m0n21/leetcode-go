/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/merge-strings-alternately/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/2 10:53
 */

package _merge_strings_alternately

func mergeAlternately(word1 string, word2 string) string {
	var ret []byte

	m, n := len(word1), len(word2)
	i, j := 0, 0
	for i < m || j < n {
		if i < m && (i == j || j == n) {
			ret = append(ret, word1[i])
			i++
		} else {
			ret = append(ret, word2[j])
			j++
		}
	}

	return string(ret)
}
