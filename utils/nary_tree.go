/**
 * @Author         : s1m0n21
 * @Description    : N-ary tree util
 * @Project        : leetcode-go
 * @File           : nary_tree.go
 * @Date           : 2021/11/22 9:52 上午
 */

package utils

type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}

func (n *NTreeNode) Preorder() []int {
	return nPreorder(n)
}

func nPreorder(node *NTreeNode) []int {
	var stack []*NTreeNode
	var nums []int

	stack = append(stack, node)

	for len(stack) > 0 {
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, node.Val)

		for i := len(node.Children)-1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}

	return nums
}

func MakeNaryTree(v ...interface{}) *NTreeNode {
	nums := newInts(v...)
	if len(nums) == 0 || nums[0].nil {
		return nil
	}
	root := makeNaryTree(nums)

	return root
}

func makeNaryTree(nums []IntValue) *NTreeNode {
	if len(nums) == 0 || nums[0].nil {
		return nil
	}

	if len(nums) == 1 {
		return &NTreeNode{Val: nums[0].val, Children: nil}
	}

	var queue = make([]*NTreeNode, 0)
	root := &NTreeNode{Val: nums[0].val}
	nums = nums[2:]
	queue = append(queue, root)

	for len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]

		for _, v := range nums {
			nums = nums[1:]

			if !v.nil {
				n := &NTreeNode{Val: v.val}
				parent.Children = append(parent.Children, n)
				queue = append(queue, n)
				continue
			}

			break
		}
	}

	return root
}
