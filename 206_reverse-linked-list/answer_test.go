/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/18 12:09 下午
 */

package _reverse_linked_list

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	head := utils.MakeListNode(1, 2, 3, 4, 5)
	reversed := reverseList(head)
	t.Logf("answer = %s", reversed.String())
}
