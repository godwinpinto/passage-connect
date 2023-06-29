package util

import (
	"fmt"
	"sync"
	"time"
)

func CountdownProgressBar(cancelCounter *CancelCounter) {

	countdown := cancelCounter.Countdown()
	for i := countdown; i >= 0; i-- {
		displayProgressBar(i, countdown)
		time.Sleep(time.Second)
	}
	fmt.Println("\nCountdown complete!")
}

func displayProgressBar(secondsElapsed, totalSeconds int) {
	barWidth := 50

	filledWidth := int(float64(secondsElapsed) / float64(totalSeconds) * float64(barWidth))
	emptyWidth := barWidth - filledWidth

	loadingBar := "[" + getRepeatedString("=", filledWidth) + getRepeatedString(" ", emptyWidth) + "]"

	fmt.Printf("\r%s %d seconds remaining", loadingBar, secondsElapsed)
}

func getRepeatedString(ch string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += ch
	}
	return result
}

type CancelCounter struct {
	countdown int
	cancelled bool
	mutex     sync.Mutex
}

func NewCancelCounter(countdown int) *CancelCounter {
	return &CancelCounter{
		countdown: countdown,
		cancelled: false,
	}
}

func (c *CancelCounter) Countdown() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.countdown
}

func (c *CancelCounter) IsCancelled() bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.cancelled
}

func (c *CancelCounter) Cancel() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cancelled = true
}
