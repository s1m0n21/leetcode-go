/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/24 1:38 上午
 */

package _latest_time_by_replacing_hidden_digits

import "testing"

func TestAnswer(t *testing.T) {
	time := "2?:?0"

	t.Logf("answer = %s", maximumTime(time))
}
