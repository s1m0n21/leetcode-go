/**
 * @Author         : s1m0n21
 * @Description    : Binary tree util
 * @Project        : leetcode-go
 * @File           : binary_tree.go
 * @Date           : 2021/7/23 11:01 上午
 */

package utils

import (
	"io/ioutil"
	"os"
	"reflect"
	"strconv"

	"github.com/awalterschulze/gographviz"
	"github.com/s1m0n21/exec-go"
)

type IntValue struct {
	val int
	nil bool
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) Get(v int) *TreeNode {
	var dfs func(*TreeNode)
	var target *TreeNode
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Val == v {
			target = node
			return
		}

		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(n)

	return target
}

func (n *TreeNode) Preorder() []int {
	return preorder(n)
}

func (n *TreeNode) Inorder() []int {
	return inorder(n)
}

func (n *TreeNode) Postorder() []int {
	return postorder(n)
}

func (n *TreeNode) Graph() {
	ast, _ := gographviz.Parse([]byte("digraph G{}"))
	graph := gographviz.NewGraph()
	gographviz.Analyse(ast, graph)

	queue := make([]*TreeNode, 0)
	queue = append(queue, n)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		graph.AddNode("G", strconv.Itoa(node.Val), nil)
		if node.Left != nil {
			graph.AddEdge(strconv.Itoa(node.Val), strconv.Itoa(node.Left.Val), true, nil)
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			graph.AddEdge(strconv.Itoa(node.Val), strconv.Itoa(node.Right.Val), true, nil)
			queue = append(queue, node.Right)
		}
	}

	ioutil.WriteFile("tree.gv", []byte(graph.String()), 0666)

	cmd := gexec.NewCommand()
	cmd.ExecSync("dot", "tree.gv", "-T", "svg", "-o", "tree.svg")
	cmd.ExecSync("open", "tree.svg")

	os.Remove("tree.gv")
}

func preorder(node *TreeNode) []int {
	var stack []*TreeNode
	var nums []int

	for node != nil || len(stack) > 0 {
		for node != nil {
			nums = append(nums, node.Val)
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return nums
}

func inorder(node *TreeNode) []int {
	var stack []*TreeNode
	var nums []int

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, node.Val)
		node = node.Right
	}

	return nums
}

func postorder(node *TreeNode) []int {
	var stack []*TreeNode
	var prev *TreeNode
	var nums []int

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Right == nil || node.Right == prev {
			nums = append(nums, node.Val)
			prev = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
		}
	}

	return nums
}

func newInts(v ...interface{}) []IntValue {
	ints := make([]IntValue, len(v))
	for i, x := range v {
		if n, ok := x.(int); ok {
			ints[i] = IntValue{val: n, nil: false}
		} else {
			ints[i] = IntValue{nil: true}
		}
	}

	return ints
}

func SameTree(a, b *TreeNode) bool {
	return reflect.DeepEqual(a.Inorder(), b.Inorder())
}

func MakeBinaryTree(v ...interface{}) *TreeNode {
	nums := newInts(v...)
	if len(nums) == 0 || nums[0].nil {
		return nil
	}
	root := makeBinaryTree(nums)

	return root
}

func makeBinaryTree(nums []IntValue) *TreeNode {
	if len(nums) == 0 || nums[0].nil {
		return nil
	}

	var queue = make([]*TreeNode, 0)
	root := &TreeNode{Val: nums[0].val}
	nums = nums[1:]
	queue = append(queue, root)

	for len(nums) >= 2 {
		node := queue[0]
		queue = queue[1:]

		if !nums[0].nil {
			node.Left = &TreeNode{Val: nums[0].val}
			queue = append(queue, node.Left)
		}

		if !nums[1].nil {
			node.Right = &TreeNode{Val: nums[1].val}
			queue = append(queue, node.Right)
		}

		nums = nums[2:]
	}

	if len(nums) > 0 {
		if !nums[0].nil {
			queue[0].Left = &TreeNode{Val: nums[0].val}
		}
	}

	return root
}
