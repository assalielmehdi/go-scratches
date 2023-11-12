package leetcode

type BrowserHistory struct {
	History      []string
	CurrentIndex int
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

func ConstructorBrowserHistory(homepage string) BrowserHistory {
	return BrowserHistory{[]string{homepage}, 0}
}

func (this *BrowserHistory) Visit(url string) {
	this.History = append(this.History[:this.CurrentIndex+1], url)

	this.CurrentIndex = len(this.History) - 1
}

func (this *BrowserHistory) Back(steps int) string {
	this.CurrentIndex = max(0, this.CurrentIndex-steps)

	return this.History[this.CurrentIndex]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.CurrentIndex = min(len(this.History)-1, this.CurrentIndex+steps)

	return this.History[this.CurrentIndex]
}
