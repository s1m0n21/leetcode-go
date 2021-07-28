/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/group-anagrams/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/7 1:53 上午
 */

package _group_anagrams

func groupAnagrams(strs []string) [][]string {
	var index = make(map[[26]int][]string)
	for _, str := range strs {
		var curr [26]int
		for i := 0; i < len(str); i++ {
			curr[str[i]-'a']++
		}

		index[curr] = append(index[curr], str)
	}

	var out  = make([][]string, len(index))
	var i = 0
	for _, s := range index {
		out[i] = s
		i++
	}

	return out
}
