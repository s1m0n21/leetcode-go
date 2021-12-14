/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/12/14 7:08 PM
 */

package _serialize_and_deserialize_binary_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	codec := Constructor()

	root := utils.MakeBinaryTree(1, 2, 3, nil, nil, 4, 5, nil, nil, 10)
	s := codec.serialize(root)
	root = codec.deserialize(s)

	t.Log(s)
	t.Log(root.Preorder()) // 1,2,3,4,5,10
}
