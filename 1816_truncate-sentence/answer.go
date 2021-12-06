/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/truncate-sentence/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/12/6 4:58 PM
 */

package _truncate_sentence

func truncateSentence(s string, k int) string {
	var i, j int
	for i < k {
		if j+1 == len(s) || s[j+1] == ' ' {
			i++
		}
		j++
	}
	return s[:j]
}
