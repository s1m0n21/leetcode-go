/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/4 11:39 上午
 */

package _longest_common_prefix

import "testing"

func TestAnswer(t *testing.T) {
	strs := []string{
		"flower",
		"flow",
		"flight",
	}

	t.Logf("strs = %+v, prefix = %s", strs, longestCommonPrefix(strs))
}
