/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/7/14 12:26
 */

package _prefix_and_suffix_search

import (
	"testing"
)

func TestSolution(t *testing.T) {
	wf := Constructor([]string{"apple", "aaaaeeee"})

	idx := wf.F("a", "le")
	t.Logf("%d", idx)
}
