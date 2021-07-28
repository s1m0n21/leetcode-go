/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/18 11:55 上午
 */

package interview_10_02_group_anagrams_lcci

import "testing"

func TestAnswer(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	t.Logf("answer = %+v", groupAnagrams(strs))
}
