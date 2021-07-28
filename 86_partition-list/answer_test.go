/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/18 4:11 下午
 */

package _partition_list

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	head := utils.MakeListNode(1,4,3,2,5,2)
	x := 3
	res := partition(head, x)

	t.Logf("answer = %s", res)
}
