/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/20 17:08
 */

package _implement_trie_prefix_tree

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	type input struct {
		op int
		s  string
	}

	trie := Constructor()
	testCase := utils.NewTestCase(t, func(i input) bool {
		if i.op == 0 {
			return trie.Search(i.s)
		}
		return trie.StartsWith(i.s)
	})

	trie.Insert("apple")

	testCase.SetAndRun(input{0, "apple"}, true)
	testCase.SetAndRun(input{0, "app"}, false)
	testCase.SetAndRun(input{1, "app"}, true)

	trie.Insert("app")

	testCase.SetAndRun(input{0, "app"}, true)
	testCase.SetAndRun(input{0, "a"}, false)
	testCase.SetAndRun(input{1, "a"}, true)
}
