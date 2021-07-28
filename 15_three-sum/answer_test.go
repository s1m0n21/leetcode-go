/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/5 12:54 上午
 */

package _three_sum

import "testing"

func TestAnswer(t *testing.T) {
	t.Logf("answer = %+v", threeSum([]int{-1,0,1,2,-1,-4}))
}
