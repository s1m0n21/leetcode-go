/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/encode-and-decode-tinyurl/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/6/29 15:39
 */

package _encode_and_decode_tinyurl

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
	idx    int
	record map[string]string
}

func Constructor() Codec {
	return Codec{0, make(map[string]string)}
}

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	c.record[strconv.Itoa(c.idx)] = longUrl
	c.idx++
	return fmt.Sprintf("https://tinyurl.com/%d", c.idx-1)
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	s := strings.Split(shortUrl, "/")
	return c.record[s[len(s)-1]]
}
