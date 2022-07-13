/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/7/13 09:37
 */

package _asteroid_collision

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	tests := utils.NewTestCase(t, asteroidCollision)

	tests.SetAndRun([]int{5, 10, -5}, []int{5, 10})
	tests.SetAndRun([]int{8, -8}, []int{})
	tests.SetAndRun([]int{10, 2, -5}, []int{10})
	tests.SetAndRun([]int{-1, 1, -5}, []int{-1, -5})
}
