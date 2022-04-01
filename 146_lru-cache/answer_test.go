/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/1 11:08
 */

package _lru_cache

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	cache := Constructor(2)
	testCase := utils.NewTestCase(t, func(i int) int {
		return cache.Get(i)
	})

	cache.Put(1, 1)
	cache.Put(2, 2)
	testCase.SetAndRun(1, 1)
	cache.Put(3, 3)
	testCase.SetAndRun(2, -1)
	cache.Put(4, 4)
	testCase.SetAndRun(1, -1)
	testCase.SetAndRun(3, 3)
	testCase.SetAndRun(4, 4)
}
