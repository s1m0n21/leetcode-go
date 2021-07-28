/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/29 9:21 上午
 */

package _container_with_most_water

import "testing"

func TestAnswer(t *testing.T) {
	height := []int{1,8,6,2,5,4,8,3,7}
	t.Logf("answer = %d", maxArea(height))
}
