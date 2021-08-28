/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/27 10:01 上午
 */

package _find_median_from_data_stream

import "testing"

func TestAnswer(t *testing.T) {
	mf := Constructor()

	mf.AddNum(1)
	mf.AddNum(2)
	t.Logf("median = %f", mf.FindMedian())
	mf.AddNum(3)
	t.Logf("median = %f", mf.FindMedian())
}
