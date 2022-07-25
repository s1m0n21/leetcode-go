/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/bulls-and-cows/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/7/25 10:33
 */

package _bulls_and_cows

import "fmt"

func getHint(secret string, guess string) string {
	var bulls, cows int
	var record = make(map[byte]int)

	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
			continue
		}
		record[secret[i]]++
	}

	for i := 0; i < len(guess); i++ {
		if guess[i] != secret[i] && record[guess[i]] > 0 {
			cows++
			record[guess[i]]--
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}
