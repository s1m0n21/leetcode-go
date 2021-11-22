/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/8 8:26 下午
 */

package _binary_search_tree_iterator

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	iter := Constructor(utils.MakeBinaryTree(7, 3, 15, nil, nil, 9, 20))

	t.Logf("next = %d", iter.Next())
	t.Logf("next = %d", iter.Next())
	t.Logf("hasNext = %t", iter.HasNext())
	t.Logf("next = %d", iter.Next())
	t.Logf("hasNext = %t", iter.HasNext())
	t.Logf("next = %d", iter.Next())
	t.Logf("hasNext = %t", iter.HasNext())
	t.Logf("next = %d", iter.Next())
	t.Logf("hasNext = %t", iter.HasNext())
}
