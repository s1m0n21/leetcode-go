/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/21 12:20 上午
 */

package _middle_of_the_linked_list

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	head := utils.MakeListNode(1,2,3,4,5)
	res := middleNode(head)

	t.Logf("answer = %s", res)
}
