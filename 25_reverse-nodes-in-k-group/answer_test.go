/**
 * @Author         : s1m0n21
 * @Description    :
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/19 9:56 下午
 */

package _reverse_nodes_in_k_group

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	head := utils.MakeListNode(1,2,3,4,5)
	k := 4
	res := reverseKGroup(head, k)

	t.Logf("answer = %s", res)
}
