/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/add-binary/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/24 5:32 下午
 */

package _add_binary

func addBinary(a string, b string) string {
	var result []byte
	ia, ib := len(a)-1, len(b)-1
	carry := 0

	for ia >= 0 || ib >= 0 {
		if ia >= 0 && ib >= 0 {
			sum := int(a[ia]-'0') + int(b[ib]-'0') + carry
			carry = 0
			if sum == 3 {
				result = append(result, '1')
				carry++
			} else if sum == 2 {
				result = append(result, '0')
				carry++
			} else {
				result = append(result, byte('0'+sum))
			}
		} else if ia >= 0 {
			sum := int(a[ia]-'0') + carry
			carry = 0
			if sum == 2 {
				result = append(result, '0')
				carry++
			} else {
				result = append(result, byte('0'+sum))
			}
		} else {
			sum := int(b[ib]-'0') + carry
			carry = 0
			if sum == 2 {
				result = append(result, '0')
				carry++
			} else {
				result = append(result, byte('0'+sum))
			}
		}
		ia--
		ib--
	}

	if carry > 0 {
		result = append(result, '1')
	}

	n := len(result) - 1
	for i := 0; i < len(result)/2; i++ {
		result[i], result[n-i] = result[n-i], result[i]
	}

	return string(result)
}
