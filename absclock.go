// Reference implementation of the Absolute Clock:
// A clock that represents a single moment in time as a unique string.
package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Get the current time broken into its components.
	// Time is represented in UTC.
	now := time.Now().UTC()
	year := strconv.Itoa(now.Year())
	month := strconv.Itoa(int(now.Month()))
	day := strconv.Itoa(now.Day())
	hour := strconv.Itoa(now.Hour())
	minute := strconv.Itoa(now.Minute())
	second := strconv.Itoa(now.Second())
	millisecond := strconv.Itoa(now.Nanosecond() / 1000000)
	nanosecond := strconv.Itoa(now.Nanosecond() / 1000) // withn the current millisecond

	// Pad the values with a leading zero if needed.
	if now.Month() < 10 {
		month = "0" + month
	}
	if now.Day() < 10 {
		day = "0" + day
	}
	if now.Hour() < 10 {
		hour = "0" + hour
	}
	if now.Minute() < 10 {
		minute = "0" + minute
	}
	if now.Second() < 10 {
		second = "0" + second
	}

	if (now.Nanosecond() / 1000000) < 100 {
		millisecond = "0" + millisecond
	}

	if (now.Nanosecond() / 1000) < 100000 {
		nanosecond = "0" + nanosecond
	}

	// Build the string
	var builder strings.Builder

	// Safe to hardcode these values; they won't change before this program is long forgotten.
	builder.WriteString("E:") // Eternity.
	builder.WriteString("4:") // Phanerozoic eon.
	builder.WriteString("3:") // Cenozoic era.
	builder.WriteString("3:") // Quaternary period.
	builder.WriteString("2:") // Holocene epoch.
	builder.WriteString("3:") // Meghalayan age.

	builder.WriteString(year) // Year
	builder.WriteString(":")
	builder.WriteString(month) // Month
	builder.WriteString(":")
	builder.WriteString(day) // Day
	builder.WriteString(":")
	builder.WriteString(hour) // Hour
	builder.WriteString(":")
	builder.WriteString(minute) // Minute
	builder.WriteString(":")
	builder.WriteString(second) // Second
	builder.WriteString(":")
	builder.WriteString(millisecond) // Millisecond
	builder.WriteString(":")
	builder.WriteString(nanosecond) // Nanosecond

	absTime := builder.String()

	fmt.Println(absTime)
}
