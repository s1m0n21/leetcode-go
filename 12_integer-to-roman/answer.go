/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/integer-to-roman/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/22 2:41 下午
 */

package _integer_to_roman

var (
	thousands = []string{"", "M", "MM", "MMM"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func intToRoman(num int) string {
	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}
