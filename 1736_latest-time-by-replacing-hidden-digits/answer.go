/**
 * @Author         : s1m0n21
 * @Description    : Answer for
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/24 1:18 上午
 */

package _latest_time_by_replacing_hidden_digits

func maximumTime(time string) string {
	var res = make([]byte, 0)
	if time[0] == '?' {
		if time[1] == '?' || time[1] < '4' {
			res = append(res, '2')
		} else {
			res = append(res, '1')
		}
	} else {
		res = append(res, time[0])
	}

	if time[1] == '?' {
		if time[0] == '?' || time[0] == '2' {
			res = append(res, '3')
		} else {
			res = append(res, '9')
		}
	} else {
		res = append(res, time[1])
	}

	res = append(res, ':')

	if time[3] == '?' {
		res = append(res, '5')
	} else {
		res = append(res, time[3])
	}

	if time[4] == '?' {
		res = append(res, '9')
	} else {
		res = append(res, time[4])
	}

	return string(res)
}