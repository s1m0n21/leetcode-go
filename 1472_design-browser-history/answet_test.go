/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answet_test.go
 * @Date           : 2021/8/2 8:08 下午
 */

package _design_browser_history

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	history := Constructor("leetcode.com")

	history.Visit("google.com")
	history.Visit("facebook.com")
	history.Visit("youtube.com")
	utils.Log.Debugf("forward = %s", history.Forward(11))
	utils.Log.Debugf("back = %s", history.Back(11))
}
