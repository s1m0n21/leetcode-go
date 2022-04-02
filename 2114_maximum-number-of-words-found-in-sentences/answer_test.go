/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/2 16:02
 */

package _maximum_number_of_words_found_in_sentences

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, mostWordsFound)

	testCase.SetAndRun([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}, 6)
	testCase.SetAndRun([]string{"please wait", "continue to fight", "continue to win"}, 3)
}
