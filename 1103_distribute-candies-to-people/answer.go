/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/distribute-candies-to-people/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/12 7:26 下午
 */

package _distribute_candies_to_people

func distributeCandies(candies int, num_people int) []int {
	var ans = make([]int, num_people)
	i, n := 0, 1
	for candies > 0 {
		if i == num_people {
			i = 0
		}

		if candies < n {
			ans[i] += candies
			candies -= candies
		} else {
			ans[i] += n
			candies -= n
		}

		n++
		i++
	}

	return ans
}
