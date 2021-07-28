/**
 * @Author         : s1m0n21
 * @Description    :
 * @Project        : leetcode-go
 * @File           : tree_test.go
 * @Date           : 2021/7/23 11:20 上午
 */

package utils

import "testing"

func TestMakeTreeNode(t *testing.T) {
	root := MakeTreeNode(2,1,3,6,4,8,5,9,10)
	root.Graph()
}
