/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/27 3:46 下午
 */

package _path_sum

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	root := utils.MakeTreeNode(5,4,8,11,nil,13,4,7,2,nil,nil,nil,1)
	targetSum := 22

	t.Logf("answer = %t", hasPathSum(root,targetSum))
}
