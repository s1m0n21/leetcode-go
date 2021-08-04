/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/7 5:10 下午
 */

package _permutations

import "testing"

func TestAnswer(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	t.Logf("answer = %+v", permute(nums))
}
