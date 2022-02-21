/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/partition-labels/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/2/21 4:18 PM
 */

package _partition_labels

func partitionLabels(s string) []int {
	var ans []int
	var last [26]int
	var start, end int

	for i, c := range s {
		last[c-'a'] = i
	}

	for i, c := range s {
		if last[c-'a'] > end {
			end = last[c-'a']
		}
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}

	return ans
}
