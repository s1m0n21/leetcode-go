/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/7/22 15:32
 */

package _longest_repeating_character_replacement

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	type input struct {
		s string
		k int
	}

	tests := utils.NewTestCase(t, func(i input) int {
		return characterReplacement(i.s, i.k)
	})

	tests.SetAndRun(input{"ABAB", 2}, 4)
	tests.SetAndRun(input{"AABABBA", 1}, 4)
}
