/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/18 4:50 下午
 */

package _odd_even_linked_list

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	head := utils.MakeListNode(2,1,3,5,6,4,7)
	res := oddEvenList(head)

	t.Logf("answer = %s", res)
}
