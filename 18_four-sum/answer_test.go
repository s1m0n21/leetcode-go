/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/5 1:21 上午
 */

package _four_sum

import "testing"

func TestAnswer(t *testing.T) {
	t.Logf("answer = %+v", fourSum([]int{1,0,-1,0,-2,2}, 0))
}
