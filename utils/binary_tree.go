/**
 * @Author         : s1m0n21
 * @Description    : Binary tree util
 * @Project        : leetcode-go
 * @File           : binary_tree.go
 * @Date           : 2021/7/23 11:01 上午
 */

package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/s1m0n21/exec-go"

	"github.com/awalterschulze/gographviz"
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

func (n *TreeNode) Graph(name string) {
	ast, _ := gographviz.Parse([]byte(`digraph G{}`))
	graph := gographviz.NewGraph()
	_ = gographviz.Analyse(ast, graph)

	queue := make([]*TreeNode, 0)
	queue = append(queue, n)

	graph.AddAttr("G", "splines", "line")

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		placeholder := fmt.Sprintf("__%p", node)

		var distance, rightDistance int
		var target, rightTarget *TreeNode

		x := node.Left
		for x != nil {
			distance++
			target = x
			x = x.Right
		}
		x = node.Right
		for x != nil {
			rightDistance++
			rightTarget = x
			x = x.Left
		}

		if rightDistance < distance {
			distance = rightDistance
			target = rightTarget
		}

		if distance > 0 {
			_ = graph.AddSubGraph("G", fmt.Sprintf("s_%p", node), map[string]string{"rank": "same"})
			_ = graph.AddNode(fmt.Sprintf("s_%p", node), placeholder, map[string]string{
				"group": fmt.Sprintf("_%p", node),
				"label": "\"\"",
				"width": "0",
				"style": "invis",
			})
			_ = graph.AddNode(fmt.Sprintf("s_%p", node), fmt.Sprintf("_%p", target), map[string]string{
				"group": fmt.Sprintf("_%p", target),
				"label": fmt.Sprintf("%d", target.Val),
				"shape": "circle",
			})
		} else {
			_ = graph.AddNode("G", placeholder, map[string]string{
				"group": fmt.Sprintf("_%p", node),
				"label": "\"\"",
				"width": "0",
				"style": "invis",
			})
		}

		if !graph.IsNode(fmt.Sprintf("_%p", node)) {
			_ = graph.AddNode("G", fmt.Sprintf("_%p", node), map[string]string{
				"group": fmt.Sprintf("_%p", node),
				"label": fmt.Sprintf("%d", node.Val),
				"shape": "circle",
			})
		}

		if node.Left == nil && node.Right == nil {
			continue
		}

		if node.Left != nil {
			_ = graph.AddEdge(fmt.Sprintf("_%p", node), fmt.Sprintf("_%p", node.Left), true, nil)
			queue = append(queue, node.Left)
		} else {
			_ = graph.AddEdge(fmt.Sprintf("_%p", node), placeholder, true, map[string]string{
				"style": "invis",
			})
		}

		_ = graph.AddEdge(fmt.Sprintf("_%p", node), placeholder, true, map[string]string{
			"group": fmt.Sprintf("_%p", node),
			"label": "\"\"",
			"width": "0",
			"style": "invis",
		})

		if node.Right != nil {
			_ = graph.AddEdge(fmt.Sprintf("_%p", node), fmt.Sprintf("_%p", node.Right), true, nil)
			queue = append(queue, node.Right)
		} else {
			_ = graph.AddEdge(fmt.Sprintf("_%p", node), placeholder, true, map[string]string{
				"style": "invis",
			})
		}
	}

	_ = ioutil.WriteFile(fmt.Sprintf("%s.gv", name), []byte(graph.String()), 0o666)

	cmd := gexec.NewCommand()
	_ = cmd.ExecSync("dot", fmt.Sprintf("%s.gv", name), "-T", "svg", "-o", fmt.Sprintf("%s.svg", name))
	_ = cmd.ExecSync("open", fmt.Sprintf("%s.svg", name))

	_ = os.Remove(fmt.Sprintf("%s.gv", name))
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

	queue := make([]*TreeNode, 0)
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
