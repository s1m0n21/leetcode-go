/**
 * @Author         : s1m0n21
 * @Description    : Log util
 * @Project        : leetcode-go
 * @File           : log.go
 * @Date           : 2021/6/6 11:19 上午
 */

package utils

import golog "github.com/s1m0n21/go-log"

var Log = golog.New("leetcode", golog.OptLevelFromEnv())
