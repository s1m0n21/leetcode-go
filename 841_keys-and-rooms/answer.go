/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/keys-and-rooms/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/25 09:51
 */

package _keys_and_rooms

func canVisitAllRooms(rooms [][]int) bool {
	unlocked := make([]bool, len(rooms))
	unlocked[0] = true
	cnt := 1

	keys := append([]int(nil), rooms[0]...)
	for len(keys) > 0 {
		key := keys[0]
		keys = keys[1:]
		if !unlocked[key] {
			cnt++
			unlocked[key] = true
			keys = append(keys, rooms[key]...)
		}
	}

	return cnt == len(rooms)
}
