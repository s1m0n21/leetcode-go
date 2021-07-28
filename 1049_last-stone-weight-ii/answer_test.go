/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/8 10:57 上午
 */

package _last_stone_weight_ii

import "testing"

func TestAnswer(t *testing.T) {
	stones := []int{2,7,4,1,8,1}
	t.Logf("answer = %d", lastStoneWeightII(stones))
}
