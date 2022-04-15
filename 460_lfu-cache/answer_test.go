/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/15 14:00
 */

package _lfu_cache

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	lfu := Constructor(3)
	testCase := utils.NewTestCase(t, lfu.Get)

	lfu.Put(2, 2)
	lfu.Put(1, 1)
	testCase.SetAndRun(2, 2)
	testCase.SetAndRun(1, 1)
	testCase.SetAndRun(2, 2)
	lfu.Put(3, 3)
	lfu.Put(4, 4)
	testCase.SetAndRun(3, -1)
	testCase.SetAndRun(2, 2)
	testCase.SetAndRun(1, 1)
	testCase.SetAndRun(4, 4)
}
