/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/design-browser-history/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/2 7:56 下午
 */

package _design_browser_history

type BrowserHistory struct {
	history []string
	index   int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		history: []string{homepage},
		index:   0,
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.history = this.history[:this.index+1]
	this.history = append(this.history, url)
	this.index++
}

func (this *BrowserHistory) Back(steps int) string {
	for i := 0; i < steps; i++ {
		if this.index == 0 {
			return this.history[this.index]
		}

		this.index--
	}

	return this.history[this.index]
}

func (this *BrowserHistory) Forward(steps int) string {
	for i := 0; i < steps; i++ {
		if this.index+1 == len(this.history) {
			return this.history[this.index]
		}

		this.index++
	}

	return this.history[this.index]
}
