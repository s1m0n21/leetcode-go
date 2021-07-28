/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/longest-common-prefix/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/4 11:31 上午
 */

package _longest_common_prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
