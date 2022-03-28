/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/28 16:28
 */

package _convert_binary_number_in_a_linked_list_to_integer

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, getDecimalValue)

	testCase.SetAndRun(utils.MakeListNode(1, 0, 1), 5)
	testCase.SetAndRun(utils.MakeListNode(0), 0)
	testCase.SetAndRun(utils.MakeListNode(1), 1)
	testCase.SetAndRun(utils.MakeListNode(1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0), 18880)
	testCase.SetAndRun(utils.MakeListNode(0, 0), 0)
}
