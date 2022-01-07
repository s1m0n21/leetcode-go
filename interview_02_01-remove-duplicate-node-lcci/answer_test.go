/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/7 12:23 PM
 */

package interview_02_01_remove_duplicate_node_lcci

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  *utils.ListNode
		expect *utils.ListNode
	}{
		{utils.MakeListNode(1, 2, 3, 3, 2, 1), utils.MakeListNode(1, 2, 3)},
		{utils.MakeListNode(1, 1, 1, 1, 2), utils.MakeListNode(1, 2)},
		{utils.MakeListNode(), utils.MakeListNode()},
	}

	for i, test := range tests {
		if actual := removeDuplicateNodes(test.input); !utils.SameList(actual, test.expect) {
			t.Errorf("%d: input = %s, exepct = %s, actual = %s", i, test.input, test.expect, actual)
		}
	}
}
