/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/11 7:57 下午
 */

package _number_of_boomerangs

import "testing"

func TestAnswer(t *testing.T) {
	points := [][]int{{0,0},{1,0},{2,0}}
	t.Logf("answer = %d", numberOfBoomerangs(points))
}
