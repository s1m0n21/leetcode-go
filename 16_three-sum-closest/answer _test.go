/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer _test.go
 * @Date           : 2021/7/5 5:53 下午
 */

package _three_sum_closest

import "testing"

func TestAnswer(t *testing.T) {
	t.Logf("answer = %d", threeSumClosest([]int{-1,0,1,1,55}, 3))
}
