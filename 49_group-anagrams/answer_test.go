/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/7 2:04 上午
 */

package _group_anagrams

import "testing"

func TestAnswer(t *testing.T) {
	t.Logf("answer = %+v", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
