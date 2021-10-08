/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/repeated-dna-sequences/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/8 9:37 上午
 */

package _repeated_dna_sequences

func findRepeatedDnaSequences(s string) []string {
	if len(s) == 0 {
		return nil
	}

	var res []string
	dna := make(map[string]int)
	for i := 0; i < len(s)-9; i++ {
		if n, has := dna[s[i:i+10]]; has && n != -1 {
			dna[s[i:i+10]] = -1
			res = append(res, s[i:i+10])
		} else if !has {
			dna[s[i:i+10]] = 1
		}
	}

	return res
}
