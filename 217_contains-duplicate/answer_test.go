/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/17 11:34 下午
 */

package _contains_duplicate

import "testing"

func TestAnswer(t *testing.T) {
	nums := []int{1,2,3,1}

	t.Logf("answer = %t", containsDuplicate(nums))
}
