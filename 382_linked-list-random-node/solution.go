/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/linked-list-random-node/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/4/26 17:53
 */

package _linked_list_random_node

import (
	"math/rand"

	"github.com/s1m0n21/leetcode-go/utils"
)

type Solution struct {
	nodes []*utils.ListNode
}

func Constructor(head *utils.ListNode) Solution {
	nodes := make([]*utils.ListNode, 0)
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	return Solution{nodes}
}

func (s *Solution) GetRandom() int {
	return s.nodes[rand.Intn(len(s.nodes))].Val
}
