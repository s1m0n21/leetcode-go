/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/12/6 11:31 AM
 */

package _insert_delete_getrandom_o1

import "testing"

func TestAnswer(t *testing.T) {
	set := Constructor()

	t.Logf("insert 1: %t, true", set.Insert(1))
	t.Logf("remove 2: %t, false", set.Remove(2))
	t.Logf("insert 2: %t, ture", set.Insert(2))
	t.Logf("insert 3: %t, ture", set.Insert(3))
	t.Logf("insert 4: %t, ture", set.Insert(4))
	t.Logf("getRandom: %d", set.GetRandom())
	t.Logf("remove 1: %t, true", set.Remove(1))
	t.Logf("insert 2: %t, false", set.Insert(2))
	t.Logf("getRandom: %d", set.GetRandom())
	t.Logf("getRandom: %d", set.GetRandom())
	t.Logf("getRandom: %d", set.GetRandom())
	t.Logf("getRandom: %d", set.GetRandom())
	t.Logf("getRandom: %d", set.GetRandom())
	t.Logf("getRandom: %d", set.GetRandom())
	t.Logf("getRandom: %d", set.GetRandom())
}
