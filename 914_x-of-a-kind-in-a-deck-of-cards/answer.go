/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/12/6 4:00 PM
 */

package _x_of_a_kind_in_a_deck_of_cards

func hasGroupsSizeX(deck []int) bool {
	if len(deck) == 0 {
		return false
	}

	m := make(map[int]int)
	for _, n := range deck {
		m[n]++
	}

	max := m[deck[0]]
	for _, n := range m {
		max = gcd(n, max)
		if max < 2 {
			return false
		}
	}

	return true
}

func gcd(x, y int) int {
	t := x % y
	if t > 0 {
		return gcd(y, t)
	}
	return y
}
