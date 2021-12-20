/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/12/16 1:38 PM
 */

package _all_nodes_distance_k_in_binary_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		root, target *utils.TreeNode
		k            int
	}

	tree1 := utils.MakeBinaryTree(3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4)

	tests := []struct {
		input  input
		expect []int
	}{
		{input{tree1, tree1.Get(5), 2}, []int{7, 4, 1}},
	}

	for i, test := range tests {
		if actual := distanceK(test.input.root, test.input.target, test.input.k); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
