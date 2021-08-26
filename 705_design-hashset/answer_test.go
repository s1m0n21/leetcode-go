/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/23 9:17 上午
 */

package _design_hashset

import "testing"

func TestAnswer(t *testing.T) {
	set := Constructor()

	set.Add(1)
	set.Add(2)
	t.Logf("1: %t", set.Contains(1))
	t.Logf("3: %t", set.Contains(3))
	set.Add(2)
	t.Logf("2: %t", set.Contains(2))
	set.Remove(2)
	t.Logf("2: %t", set.Contains(2))
}
