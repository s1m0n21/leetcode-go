/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/9/9 9:22 上午
 */

package _design_hashmap

import "testing"

func TestAnswer(t *testing.T) {
	m := Constructor()

	m.Put(1, 1)
	m.Put(2, 2)
	t.Log(m.Get(1))
	t.Log(m.Get(3))
	m.Put(2, 1)
	t.Log(m.Get(2))
	m.Remove(2)
	t.Log(m.Get(2))
}
