/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/22 11:47 上午
 */

package _shuffle_an_array

import "testing"

func TestAnswer(t *testing.T) {
	rand := Constructor([]int{1, 2, 3})

	t.Logf("%+v", rand.Shuffle())
	t.Logf("%+v", rand.Shuffle())
	t.Logf("%+v (reset)", rand.Reset())
	t.Logf("%+v", rand.Shuffle())
	t.Logf("%+v", rand.Shuffle())
	t.Logf("%+v (reset)", rand.Reset())
}
