/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/19 11:14 下午
 */

package _insertion_sort_list

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	head := utils.MakeListNode(4,3,2,1)
	res := insertionSortList(head)

	t.Logf("answer = %s", res)
}
