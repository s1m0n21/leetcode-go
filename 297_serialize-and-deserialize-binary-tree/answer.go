/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/12/14 6:56 PM
 */

package _serialize_and_deserialize_binary_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"strconv"
	"strings"
)

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *utils.TreeNode) string {
	if root == nil {
		return ""
	}

	var b []byte
	var queue []*utils.TreeNode

	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			b = append(b, 'n')
			if len(queue) > 0 {
				b = append(b, ',')
			}
			continue
		}

		b = append(b, []byte(strconv.Itoa(node.Val))...)
		b = append(b, ',')

		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}

	return string(b)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *utils.TreeNode {
	if data == "" {
		return nil
	}

	var queue []*utils.TreeNode

	nodes := strings.Split(data, ",")
	n, _ := strconv.Atoi(nodes[0])
	root := &utils.TreeNode{Val: n}
	nodes = nodes[1:]
	queue = append(queue, root)

	for len(nodes) >= 2 {
		node := queue[0]
		queue = queue[1:]
		if nodes[0] != "n" {
			n, _ := strconv.Atoi(nodes[0])
			node.Left = &utils.TreeNode{Val: n}
			queue = append(queue, node.Left)
		}
		if nodes[1] != "n" {
			n, _ := strconv.Atoi(nodes[1])
			node.Right = &utils.TreeNode{Val: n}
			queue = append(queue, node.Right)
		}
		nodes = nodes[2:]
	}

	if len(nodes) > 0 {
		if nodes[0] != "n" {
			n, _ := strconv.Atoi(nodes[0])
			queue[0].Left = &utils.TreeNode{Val: n}
		}
	}

	return root
}
