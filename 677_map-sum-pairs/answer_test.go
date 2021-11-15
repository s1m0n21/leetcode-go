/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/15 10:00 上午
 */

package _map_sum_pairs

import "testing"

func TestAnswer(t *testing.T) {
	mp := Constructor()

	mp.Insert("apple", 3)
	t.Logf("sum = %d", mp.Sum("ap")) // 3
	mp.Insert("app", 2)
	t.Logf("sum = %d", mp.Sum("ap")) // 5
	mp.Insert("app", 3)
	t.Logf("sum = %d", mp.Sum("ap")) // 6
	mp.Insert("a", 3)
	t.Logf("sum = %d", mp.Sum("ap")) // 6
	t.Logf("sum = %d", mp.Sum("a"))  // 9
}
