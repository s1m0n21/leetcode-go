/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/27 4:07 下午
 */

package _sum_root_to_leaf_numbers

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	root := utils.MakeTreeNode(4,9,0,5,1)

	t.Logf("answer = %d", sumNumbers(root))
}
