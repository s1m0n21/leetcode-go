/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/18 1:57 下午
 */

package _reverse_linked_list_ii

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	head := utils.MakeListNode(1, 2, 3, 4, 5, 6)
	left := 2
	right := 4
	reversed := reverseBetween(head, left, right)

	t.Logf("answer = %s", reversed.String())
}
