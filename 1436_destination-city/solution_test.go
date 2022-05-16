/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/16 17:47
 */

package _destination_city

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	testCase := utils.NewTestCase(t, destCity)

	testCase.SetAndRun([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}, "Sao Paulo")
	testCase.SetAndRun([][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}, "A")
	testCase.SetAndRun([][]string{{"A", "Z"}}, "Z")
}
