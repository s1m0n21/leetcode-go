/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/30 2:45 下午
 */

package _minimum_window_substring

import "testing"

func TestAnswer(t *testing.T) {
	s := "cgklivwehljxrdzpfdqsapogwvjtvbzahjnsejwnuhmomlfsrvmrnczjzjevkdvroiluthhpqtffhlzyglrvorgnalk"
	p := "mqfff"

	t.Logf("answer = %s", minWindow(s, p))
}
