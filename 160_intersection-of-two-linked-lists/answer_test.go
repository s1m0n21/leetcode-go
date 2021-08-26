/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/26 9:50 上午
 */

package _intersection_of_two_linked_lists

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	headA := utils.MakeListNode(91,92,93,1,2)
	headB := utils.MakeListNode(81,82)

	tempA := headA
	for i := 0; i < 3; i++ {
		tempA = tempA.Next
	}

	tempB := headB
	for tempB.Next != nil {
		tempB = tempB.Next
	}

	tempB.Next = tempA

	res := getIntersectionNode(headA, headB)

	t.Logf("headA = %s", headA)
	t.Logf("headB = %s", headB)
	t.Logf("answer = %s", res)
}
