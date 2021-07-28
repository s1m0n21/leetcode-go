/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/28 3:15 下午
 */

package _sum_of_left_leaves

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	root := utils.MakeTreeNode(3,9,20,nil,nil,15,7)

	t.Logf("answer = %d", sumOfLeftLeaves(root))
}
