/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/longest-substring-of-all-vowels-in-order/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/13 3:55 下午
 */

package _longest_substring_of_all_vowels_in_order

var order = map[byte]int{
	'a': 0, 'e': 1, 'i': 2, 'o': 3, 'u': 4,
}

func longestBeautifulSubstring(word string) int {
	if len(word) < 5 {
		return 0
	}

	l, r, count, ans := 0, 1, 1, 0
	for ; r < len(word); r++ {
		if word[r] != word[r-1] {
			count++
			if order[word[r]] - order[word[r-1]] != 1 {
				l = r
				count = 1
				continue
			}
		}

		if count == 5 && r-l+1 > ans {
			ans = r-l+1
		}
	}

	return ans
}
