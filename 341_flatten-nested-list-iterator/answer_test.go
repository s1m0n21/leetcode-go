/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/23 3:51 下午
 */

package _flatten_nested_list_iterator

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	list := []*utils.NestedInteger{
		utils.MakeNestedInteger(1),
		utils.MakeNestedInteger(2),
		utils.MakeNestedInteger(3),
		utils.MakeNestedIntegerList(4, 5, 6),
		utils.MakeNestedIntegerList(7,8,9).Add(utils.MakeNestedIntegerList(0,1,2)),
	}

	iter := Constructor(list)
	for iter.HasNext() {
		t.Logf("got = %d", iter.Next())
	}
}
