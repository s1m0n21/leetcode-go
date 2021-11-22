/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/26 4:24 下午
 */

package _same_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	p := utils.MakeBinaryTree(1,2,3)
	q := utils.MakeBinaryTree(1,2,3)

	t.Logf("answer = %t", isSameTree(p, q))
}
