/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/18 3:42 下午
 */

package _remove_duplicates_from_sorted_list

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	head := utils.MakeListNode(1, 1, 2, 3, 3)
	res := deleteDuplicates(head)

	t.Logf("answer = %s", res)
}
