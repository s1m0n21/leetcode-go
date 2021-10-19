/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/19 4:31 下午
 */

package _design_add_and_search_words_data_structure

import "testing"

func TestAnswer(t *testing.T) {
	w := Constructor()

	w.AddWord("bad")
	w.AddWord("dad")
	w.AddWord("mad")

	t.Logf("%t", w.Search("pad")) // false
	t.Logf("%t", w.Search("bad")) // true
	t.Logf("%t", w.Search(".ad")) // true
	t.Logf("%t", w.Search("b..")) // true
	t.Logf("%t", w.Search("...")) // true
	t.Logf("%t", w.Search("ba"))  // false
}
