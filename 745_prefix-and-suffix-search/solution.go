/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/prefix-and-suffix-search/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/7/14 11:56
 */

package _prefix_and_suffix_search

type dictTree struct {
	next    [26]*dictTree
	indices []int
}

type WordFilter struct {
	prefix *dictTree
	suffix *dictTree
}

func Constructor(words []string) WordFilter {
	prefix := &dictTree{}
	suffix := &dictTree{}
	for i, word := range words {
		currPrefix := prefix
		for j := 0; j < len(word); j++ {
			if currPrefix.next[word[j]-'a'] == nil {
				currPrefix.next[word[j]-'a'] = &dictTree{indices: []int{i}}
				currPrefix = currPrefix.next[word[j]-'a']
			} else {
				currPrefix = currPrefix.next[word[j]-'a']
				currPrefix.indices = append(currPrefix.indices, i)
			}
		}

		currSuffix := suffix
		for j := len(word)-1; j >= 0; j-- {
			if currSuffix.next[word[j]-'a'] == nil {
				currSuffix.next[word[j]-'a'] = &dictTree{indices: []int{i}}
				currSuffix = currSuffix.next[word[j]-'a']
			} else {
				currSuffix = currSuffix.next[word[j]-'a']
				currSuffix.indices = append(currSuffix.indices, i)
			}
		}
	}

	return WordFilter{
		prefix: prefix,
		suffix: suffix,
	}
}

func (wf *WordFilter) F(pref string, suff string) int {
	var prefixIndices, suffixIndices []int

	currPrefix := wf.prefix
	currSuffix := wf.suffix

	for i := 0; i < len(pref); i++ {
		if currPrefix.next[pref[i]-'a'] == nil {
			return -1
		}
		currPrefix = currPrefix.next[pref[i]-'a']
	}
	prefixIndices = currPrefix.indices

	for i := len(suff)-1; i >= 0; i-- {
		if currSuffix.next[suff[i]-'a'] == nil {
			return -1
		}
		currSuffix = currSuffix.next[suff[i]-'a']
	}
	suffixIndices = currSuffix.indices

	if len(prefixIndices) == 0 || len(suffixIndices) == 0 {
		return -1
	}

	i := len(prefixIndices)-1
	j := len(suffixIndices)-1
	for i >= 0 && j >= 0 {
		if prefixIndices[i] == suffixIndices[j] {
			return prefixIndices[i]
		} else if prefixIndices[i] > suffixIndices[j] {
			i--
		} else {
			j--
		}
	}

	return -1
}
