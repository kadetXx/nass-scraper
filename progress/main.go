// reference: https://github.com/raulfilimon/golang-progressbar

package progress

import (
	"fmt"
	"time"
)

type ProgressBar struct {
	Total     int
	Current   int
	BarWidth  int
	StartTime time.Time
}

func NewProgressBar(total int, barWidth int) *ProgressBar {
	return &ProgressBar{
		Total:     total,
		Current:   0,
		BarWidth:  barWidth,
		StartTime: time.Now(),
	}
}

func (p *ProgressBar) Increment() {
	p.Current++
	p.Render()
}

func (p *ProgressBar) Render() {
	progress := (p.Current * p.BarWidth) / p.Total
	bar := "[" + stringOfChar(progress, "=") + stringOfChar(p.BarWidth-progress, " ") + "]"
	elapsed := time.Since(p.StartTime)

	fmt.Printf("\rScraping %s %d/%d Time elapsed: %s", bar, p.Current, p.Total, elapsed.String())
}

func stringOfChar(n int, char string) string {
	var result string

	for i := 0; i < n; i++ {
		result += char
	}

	return result
}
