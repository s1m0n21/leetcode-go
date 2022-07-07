/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/7/7 11:58
 */

package _parse_lisp_expression

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	tests := utils.NewTestCase(t, evaluate)

	tests.SetAndRun("(let x 2 (mult x (let x 3 y 4 (add x y))))", 14)
	tests.SetAndRun("(let x 3 x 2 x)", 2)
	tests.SetAndRun("(let x 1 y 2 x (add x y) (add x y))", 5)
	tests.SetAndRun("(let x 2 (add (let x 3 (let x 4 x)) x))", 6)
}
