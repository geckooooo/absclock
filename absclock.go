package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Get the current time broken into its components.
	// Time is represented in Zulu time.
	now := time.Now().UTC()
	year := strconv.Itoa(now.Year())
	month := strconv.Itoa(int(now.Month()))
	day := strconv.Itoa(now.Day())
	hour := strconv.Itoa(now.Hour())
	minute := strconv.Itoa(now.Minute())
	second := strconv.Itoa(now.Second())
	millisecond := strconv.Itoa(now.Nanosecond() / 1000000)
	nanosecond := strconv.Itoa(now.Nanosecond() / 1000) // withn the current millisecond

	// Pad the values with a leading zero if tanyhey are less than 10.
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

	// Safe to hardcode these values.
	builder.WriteString("E:")  // Eternity.
	builder.WriteString("4:")  // Phanerozoic eon.
	builder.WriteString("10:") // Cenozoic era.
	builder.WriteString("??:") // Quaternary period.
	builder.WriteString("??:") // Holocene epoch.
	builder.WriteString("??:") // Meghalayan age.

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
