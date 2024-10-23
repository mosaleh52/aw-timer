package helpers

import (
	"fmt"
	"strconv"
	"time"
)

func TimeDifferenceFromNow(timestamp string) string {
	// TODO:move layout as a prameter
	layout := "2006-01-02T15:04:05.999999-07:00"

	t, err := time.Parse(layout, timestamp)
	if err != nil {
		return fmt.Sprintf("Error parsing timestamp: %v", err)
	}

	now := time.Now()

	duration := t.Sub(now)

	return formatDuration(duration)
}

func formatDuration(d time.Duration) string {
	// days := int(d.Hours() / 24)
	// hours := int(d.Hours()) % 24
	minutes := int(d.Minutes())
	// seconds := int(d.Seconds()) % 60

	// return fmt.Sprintf("%d days, %d hours, %d minutes, %d seconds", days, hours, minutes, seconds)
	return strconv.Itoa(minutes * -1)
}
