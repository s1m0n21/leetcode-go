/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/26 9:55 下午
 */

package _sort_colors

import "testing"

func TestAnswer(t *testing.T) {
	nums := []int{2,0,2,1,1,0}
	sortColors(nums)
	t.Logf("answer = %+v", nums)
}
