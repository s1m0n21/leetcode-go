/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/20 9:38 下午
 */

package _rotate_list

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	head := utils.MakeListNode(1,2,3)
	k := 2000000000
	res := rotateRight(head, k)

	t.Logf("answer = %s", res)
}
