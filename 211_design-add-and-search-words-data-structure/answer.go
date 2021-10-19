/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/design-add-and-search-words-data-structure/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/19 4:23 下午
 */

package _design_add_and_search_words_data_structure

type WordDictionary struct {
	words map[int][]string
}

func Constructor() WordDictionary {
	return WordDictionary{make(map[int][]string)}
}

func (this *WordDictionary) AddWord(word string) {
	n := len(word)
	this.words[n] = append(this.words[n], word)
}

func (this *WordDictionary) Search(word string) bool {
	n := len(word)
	if words, has := this.words[n]; has {
		for _, w := range words {
			match := true
			for i := 0; i < n; i++ {
				if word[i] == '.' {
					continue
				} else if word[i] != w[i] {
					match = false
					break
				}
			}

			if match {
				return true
			}
		}
	}

	return false
}
