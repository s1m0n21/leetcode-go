/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/2 4:54 下午
 */

package _letter_combinations_of_a_phone_number

var letters = []string{
	"",
	"",
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

func findCombination(digits, s string, index int) []string {
	if index == len(digits) {
		return []string{s}
	}

	letter := letters[digits[index]-'0']
	var combinations []string
	for i := 0; i < len(letter); i++ {
		combinations = append(combinations, findCombination(digits, s+string(letter[i]), index+1)...)
	}

	return combinations
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	return findCombination(digits, "", 0)
}
