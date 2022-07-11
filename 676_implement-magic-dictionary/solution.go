/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/implement-magic-dictionary/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/7/11 10:16
 */

package _implement_magic_dictionary

type MagicDictionary struct {
	words map[int][]string
}

func Constructor() MagicDictionary {
	return MagicDictionary{make(map[int][]string)}
}

func (m *MagicDictionary) BuildDict(dictionary []string) {
	for _, str := range dictionary {
		m.words[len(str)] = append(m.words[len(str)], str)
	}
}

func (m *MagicDictionary) Search(searchWord string) bool {
	n := len(searchWord)
	if strs, in := m.words[n]; in {
		for _, str := range strs {
			var diff int
			for i := 0; i < n; i++ {
				if str[i] != searchWord[i] {
					diff++
				}
				if diff > 1 {
					break
				}
			}
			if diff == 1 {
				return true
			}
		}
	}

	return false
}
