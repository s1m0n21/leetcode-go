/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/6 11:36 下午
 */

package _linked_list_cycle_ii

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	head := utils.MakeCycleListNode(1, 2, 3, 4, 5)
	t.Logf("answer = %#v", detectCycle(head))
}
