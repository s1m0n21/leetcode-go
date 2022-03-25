/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/25 10:00
 */

package _keys_and_rooms

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, canVisitAllRooms)

	testCase.SetAndRun([][]int{{1}, {2}, {3}, {}}, true)
	testCase.SetAndRun([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false)
	testCase.SetAndRun([][]int{{2}, {}, {1}}, true)
}
