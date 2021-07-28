/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/group-anagrams-lcci/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/18 11:48 上午
 */

package interview_10_02_group_anagrams_lcci

import "sort"

func groupAnagrams(strs []string) [][]string {
	var index = make(map[[26]int][]string)
	for _, str := range strs {
		var curr [26]int
		for i := 0; i < len(str); i++ {
			curr[str[i]-'a']++
		}

		index[curr] = append(index[curr], str)
	}

	var res = make([][]string, len(index))
	var i = 0
	for _, s := range index {
		res[i] = s
		sort.Strings(s)
		i++
	}

	return res
}
