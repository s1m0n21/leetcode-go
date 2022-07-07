/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/parse-lisp-expression/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/7/7 10:36
 */

package _parse_lisp_expression

import (
	"strconv"
)

func evaluate(expression string) int {
	return processExpr(expression, make(map[string]int))
}

func getVarValue(v string, vars map[string]int) int {
	if n, in := vars[v]; in {
		return n
	}
	n, _ := strconv.Atoi(v)
	return n
}

func processExpr(expr string, outsideVars map[string]int) int {
	vars := make(map[string]int)
	for k, v := range outsideVars {
		vars[k] = v
	}

	if expr[0] != '(' {
		return getVarValue(expr, vars)
	}

	var n int
	expr = expr[1 : len(expr)-1]
	tokens := split(expr)

	if tokens[0] == "let" {
		for i := 1; i < len(tokens)-1; i += 2 {
			vars[tokens[i]] = processExpr(tokens[i+1], vars)
		}
		n = processExpr(tokens[len(tokens)-1], vars)
	}
	if tokens[0] == "mult" {
		n = processExpr(tokens[1], vars) * processExpr(tokens[2], vars)
	}
	if tokens[0] == "add" {
		n = processExpr(tokens[1], vars) + processExpr(tokens[2], vars)
	}

	return n
}
func split(s string) []string {
	var strs []string

	i, j := 0, 0
	leftBrackets, rightBrackets := 0, 0
	for ; j < len(s); j++ {
		if s[j] == '(' {
			leftBrackets++
		}
		if s[j] == ')' {
			rightBrackets++
		}
		if s[j] == ' ' && leftBrackets == rightBrackets {
			strs = append(strs, s[i:j])
			i = j + 1
			leftBrackets = 0
			rightBrackets = 0
		}
	}
	strs = append(strs, s[i:j])

	return strs
}
