/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/25 1:47 下午
 */

package _move_zeros

import "testing"

func TestAnswer(t *testing.T) {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	t.Logf("answer = %+v", nums)
}
