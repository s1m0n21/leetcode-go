/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/4/26 17:56
 */

package _linked_list_random_node

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	freq := make(map[int]int)
	count := 1000
	s := Constructor(utils.MakeListNode(1, 2, 3, 4))

	for i := 0; i < count; i++ {
		n := s.GetRandom()
		freq[n]++
	}

	for k, v := range freq {
		t.Logf("%d: %.1f%%", k, float64(v)/float64(count)*100)
	}
}
