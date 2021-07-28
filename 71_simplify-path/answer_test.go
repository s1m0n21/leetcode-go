/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/23 12:20 上午
 */

package _simplify_path

import "testing"

func TestAnswer(t *testing.T) {
	path := "/../"

	t.Logf("answer = %s", simplifyPath(path))
}
