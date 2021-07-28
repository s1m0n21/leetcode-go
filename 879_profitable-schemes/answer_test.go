/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/9 7:18 下午
 */

package _profitable_schemes

import "testing"

func TestAnswer(t *testing.T) {
	t.Logf("answer = %d", profitableSchemes(5, 3, []int{2, 2}, []int{2, 3}))
}
