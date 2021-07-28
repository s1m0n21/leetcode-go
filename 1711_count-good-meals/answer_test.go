/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/7 5:29 下午
 */

package _count_good_meals

import "testing"

func TestAnswer(t *testing.T) {
	t.Logf("answer = %d", countPairs([]int{1,3,5,7,9}))
}
