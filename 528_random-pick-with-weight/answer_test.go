/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/30 10:04 上午
 */

package _random_pick_with_weight

import "testing"

func TestAnswer(t *testing.T) {
	weights := []int{1, 2, 3, 4, 5}
	selector := Constructor(weights)

	res := make([]int, len(weights))
	for i := 0; i < 1000; i++ {
		res[selector.PickIndex()]++
	}

	t.Logf("%+v", res)
}
