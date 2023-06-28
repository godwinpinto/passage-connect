package util

import (
	"fmt"
	"time"
)

func CountdownProgressBar(countdown int) {

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
