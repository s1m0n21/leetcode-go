/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/1 17:26
 */

package _find_the_highest_altitude

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, largestAltitude)

	testCase.SetAndRun([]int{-5, 1, 5, 0, -7}, 1)
	testCase.SetAndRun([]int{-4, -3, -2, -1, 4, 3, 2}, 0)
}
